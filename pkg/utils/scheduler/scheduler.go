package scheduler

import (
	"context"
	"fmt"
	"hash/fnv"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type SchedulerOption struct {
	// taskRegister is the register of task
	// The key is the unique identifier of the task
	// The value is the flag to indicate whether the task is running
	// 任务注册表，键是任务的唯一标识符，值是指示任务是否正在运行的标志
	LogPrefix string
	LogSuffix string
	LogFunc   func(...any)

	// Silent is the flag to control whether to print log
	// 静默标志
	Silent bool

	// Debug is the flag to control whether to print debug log
	// 调试标志
	Debug bool

	// NShard is the number of shard to resolve the shortcoming of single mutex lock
	// 分片数
	NShard int
}

type Scheduler struct {
	// taskRegister is the register of task
	// The key is the unique identifier of the task
	// The value is the flag to indicate whether the task is running
	// 任务注册表，键是任务的唯一标识符，值是指示任务是否正在运行的标志
	taskRegister sync.Map
	// partitionMutex is the mutex of partition
	// 分区互斥锁
	partitionMutex []*sync.Mutex

	// logPrefix is the prefix of log
	// logSuffix is the suffix of log
	// log is the log function
	// 日志前缀、日志后缀、日志函数
	logPrefix string
	logSuffix string
	log       func(...any)

	// debug is the flag to control whether to print debug log
	// 调试标志
	debug bool

	// nshard is the number of shard to resolve the shortcoming of single mutex lock
	// 分片数
	nshard int

	// ctx is the context of scheduler
	// cancle is the cancel function of scheduler
	// 调度器的上下文、调度器的取消函数
	ctx    context.Context
	cancle context.CancelFunc
}

// NewScheduler creates a new scheduler with the given option
// 创建调度器，选项是调度器的配置
func NewScheduler(option SchedulerOption) *Scheduler {
	if option.LogFunc == nil {
		option.LogFunc = log.Println
	}
	if option.Silent {
		option.LogFunc = func(...any) {}
	}
	if option.NShard <= 32 {
		option.NShard = 32
	}
	partitionMutex := make([]*sync.Mutex, option.NShard)
	for i := 0; i < option.NShard; i++ {
		partitionMutex[i] = &sync.Mutex{}
	}
	ctx, cancle := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	s := Scheduler{
		taskRegister:   sync.Map{},
		partitionMutex: partitionMutex,

		logPrefix: option.LogPrefix,
		logSuffix: option.LogSuffix,
		log:       option.LogFunc,

		debug:  option.Debug,
		nshard: option.NShard,

		ctx:    ctx,
		cancle: cancle,
	}
	go func() {
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		s.log = func(a ...any) {}
		option.LogFunc(option.LogPrefix, "Gracefully shutting down the scheduler, please wait...", option.LogSuffix)
		cancle()
		signal.Stop(c)
		close(c)
	}()
	return &s
}

// ihash is the hash function for string
// It is used to hash the key to shard
// 字符串哈希函数，用于将键哈希到分片
func ihash(s string) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32())
}

// Start a task with the given key, offset, function and channel
// The key is the unique identifier of the task
// The offset is the delay time of the task
// The function is the task to be executed
// The channel is the channel to return the result of the task whether it is accepted
// The task is executed asynchronously
// Different from StartWithReturn, the task result is returned asynchronously
// And the task is accepted asynchronously
// 延迟任务异步调度，异步返回是否接受任务（与同步返回的区别在于，启动任务是异步的）
func (s *Scheduler) Start(key string, offset time.Duration, f func(), c ...chan<- bool) {
	if s.debug {
		s.log(s.logPrefix, fmt.Sprintf("[Key: %s]", key), " is started", s.logSuffix)
	}
	if len(c) > 1 {
		s.log(s.logPrefix, fmt.Sprintf("[Key: %s]", key), " has more than one channel", s.logSuffix)
	}
	if len(c) > 0 {
		go func() { c[0] <- s.StartWithReturn(key, offset, f) }()
	} else {
		go s.StartWithReturn(key, offset, f)
	}
}

// StartWithReturn a task with the given key, offset and function
// The key is the unique identifier of the task
// The offset is the delay time of the task
// The function is the task to be executed
// The task is executed asynchronously
// The function returns whether the task is accepted
// Different from Start, the task result is returned synchronously
// And the task is accepted synchronously
// 延迟任务异步调度，同步返回是否接受任务（与异步返回的区别在于，启动任务是顺序的）
func (s *Scheduler) StartWithReturn(key string, offset time.Duration, f func()) bool {
	shard := ihash(key) % s.nshard
	s.partitionMutex[shard].Lock()
	defer s.partitionMutex[shard].Unlock()
	if _, existed := s.taskRegister.Load(key); existed {
		if s.debug {
			s.log(s.logPrefix, fmt.Sprintf("[Shard: %d, Key: %s]", shard, key), " is already running[rejected by existed key]", s.logSuffix)
		}
		return false
	}
	ctx, cancel := context.WithCancel(context.Background())
	s.taskRegister.Store(key, cancel)
	go func(gshard int, signal *context.Context) {
		ticker := time.NewTicker(offset)
		select {
		case <-s.ctx.Done():
			ticker.Stop()
			return
		case <-(*signal).Done():
			if s.debug {
				s.log(s.logPrefix, fmt.Sprintf("[Shard: %d, Key: %s]", shard, key), " is stopped", s.logSuffix)
			}
			ticker.Stop()
			return
		case <-ticker.C:
			f()
			s.partitionMutex[gshard].Lock()
			s.taskRegister.Delete(key)
			s.partitionMutex[gshard].Unlock()
			s.log(s.logPrefix, fmt.Sprintf("[Shard: %d, Key: %s]", shard, key), " is done", s.logSuffix)
			ticker.Stop()
			return
		}
	}(shard, &ctx)
	return true
}

// ForceStartTask starts a task with the given key, offset, function and channel
// It tries to stop the existed task with the given key, and then start the given task
// The key is the unique identifier of the task
// The offset is the delay time of the task
// The function is the task to be executed
// The channel is the channel to return the result of the task whether it is accepted
// The task is executed asynchronously
// Different from ForceStartTaskWithReturn, the task result is returned asynchronously
// And the task is accepted asynchronously
// 强制覆盖任务异步调度，异步返回是否接受任务（与同步返回的区别在于，启动任务是异步的）
func (s *Scheduler) ForceStartTask(key string, offset time.Duration, f func(), c ...chan<- bool) {
	if s.debug {
		s.log(s.logPrefix, fmt.Sprintf("[Key: %s]", key), " is forced to start", s.logSuffix)
	}
	if len(c) > 1 {
		s.log(s.logPrefix, fmt.Sprintf("[Key: %s]", key), " has more than one channel", s.logSuffix)
	}
	if len(c) > 0 {
		go func() { c[0] <- s.ForceStartTaskWithReturn(key, offset, f) }()
	} else {
		go s.ForceStartTaskWithReturn(key, offset, f)
	}
}

// ForceStartTaskWithReturn starts a task with the given key, offset and function
// It tries to stop the existed task with the given key, and then start the given task
// The key is the unique identifier of the task
// The offset is the delay time of the task
// The function is the task to be executed
// The task is executed asynchronously
// The function returns whether the task is accepted
// Different from ForceStartTask, the task result is returned synchronously
// And the task is accepted synchronously
// 强制覆盖任务异步调度，同步返回是否接受任务（与异步返回的区别在于，启动任务是顺序的）
func (s *Scheduler) ForceStartTaskWithReturn(key string, offset time.Duration, f func()) bool {
	shard := ihash(key) % s.nshard
	s.partitionMutex[shard].Lock()
	defer s.partitionMutex[shard].Unlock()
	if cancel, existed := s.taskRegister.Load(key); existed {
		if s.debug {
			s.log(s.logPrefix, fmt.Sprintf("[Shard: %d, Key: %s]", shard, key), " try to force to stop the existed task", s.logSuffix)
		}
		cancel.(context.CancelFunc)()
		s.taskRegister.Delete(key)
	}
	ctx, cancel := context.WithCancel(context.Background())
	s.taskRegister.Store(key, cancel)
	go func(gshard int, signal *context.Context) {
		ticker := time.NewTicker(offset)
		select {
		case <-s.ctx.Done():
			ticker.Stop()
			return
		case <-(*signal).Done():
			if s.debug {
				s.log(s.logPrefix, fmt.Sprintf("[Shard: %d, Key: %s]", shard, key), " is stopped", s.logSuffix)
			}
			ticker.Stop()
			return
		case <-ticker.C:
			f()
			s.partitionMutex[gshard].Lock()
			s.taskRegister.Delete(key)
			s.partitionMutex[gshard].Unlock()
			s.log(s.logPrefix, fmt.Sprintf("[Shard: %d, Key: %s]", shard, key), " is done", s.logSuffix)
			ticker.Stop()
			return
		}
	}(shard, &ctx)
	return true
}
