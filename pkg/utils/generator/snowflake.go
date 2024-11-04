package generator

import (
	"fmt"
	"sync"
	"time"
)

const (
	// epoch is the number of milliseconds since a certain comment
	// This is the epoch used for the default Snowflake
	// Epoch是从某一刻开始的毫秒数，这是默认Snowflake使用的纪元
	epoch = int64(1704067200000)

	// nodeBits is the number of bits used to represent a node number
	// nodeBits是用于表示节点编号的位数
	nodeBits = uint(10)

	// stepBits is the number of bits used to represent a step number
	// stepBits是用于表示步数的位数
	stepBits = uint(12)

	// nodeMax is the maximum node number
	// nodeMax是最大节点编号
	nodeMax = int64(-1 ^ (-1 << nodeBits))

	// stepMask is the maximum step number
	// stepMask是最大步数
	stepMask = int64(-1 ^ (-1 << stepBits))

	// nodeShift is the number of bits to shift to the left when getting the node number
	// nodeShift是获取节点编号时向左移动的位数
	nodeShift = stepBits

	// timeShift is the number of bits to shift to the left when getting the timestamp
	// timeShift是获取时间戳时向左移动的位数
	timeShift = stepBits + nodeBits
)

// Snowflake is a distributed unique ID generator
// Snowflake是一个分布式唯一ID生成器
type Snowflake struct {
	node  int64
	step  int64
	mutex sync.Mutex
	last  int64
}

// NewSnowflake creates a new Snowflake instance
// NewSnowflake创建一个新的Snowflake实例
// node is the unique number of the current node
// node是当前节点的唯一编号
// node可取值范围为0到1023
func NewSnowflake(node int64) (*Snowflake, error) {
	if node < 0 || node > nodeMax {
		return nil, fmt.Errorf("Node number must be between 0 and %d", nodeMax)
	}
	return &Snowflake{
		node: node,
	}, nil
}

// Generate generates a new unique ID
// Generate生成一个新的唯一ID
func (s *Snowflake) Generate() int64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	now := time.Now().UnixNano() / 1e6
	if s.last == now {
		s.step = (s.step + 1) & stepMask
		if s.step == 0 {
			for now <= s.last {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		s.step = 0
	}

	s.last = now

	id := (now-epoch)<<timeShift | (s.node << nodeShift) | s.step
	return id
}
