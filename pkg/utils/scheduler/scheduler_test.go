package scheduler_test

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	"sfw/pkg/utils/scheduler"
)

func TestStartNoCrash(t *testing.T) {
	keys := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}
	ticker := time.NewTicker(20 * time.Second)
	s := scheduler.NewScheduler(scheduler.SchedulerOption{Silent: true})
	wg := &sync.WaitGroup{}
	for {
		select {
		case <-ticker.C:
			// if no crash, then pass
			wg.Wait()
			t.Log("TestStartNoCrash passed")
			return
		default:
			for _, key := range keys {
				wg.Add(1)
				go s.Start(key, time.Duration(100*time.Millisecond), func() { wg.Done() })
				time.Sleep(time.Duration(time.Duration(rand.Intn(100)) * time.Millisecond))
			}
		}
	}
}

func TestStartHandledCorrectly(t *testing.T) {
	key := "1"
	s := scheduler.NewScheduler(scheduler.SchedulerOption{Debug: true, LogFunc: t.Log})
	s.StartWithReturn(key, time.Duration(1*time.Second), func() {})
	for i := 0; i < 10; i++ {
		if s.StartWithReturn(key, time.Duration(1*time.Second), func() {}) {
			t.Error("TestStartHandledCorrectly failed: Start should return false")
			return
		}
	}
	time.Sleep(2 * time.Second)
	if !s.StartWithReturn(key, time.Duration(1*time.Second), func() {}) {
		t.Error("TestStartHandledCorrectly failed: Start should return true")
		return
	}
	t.Log("TestStartHandledCorrectly passed")
}

func TestStartHandledCorrectlyWithMultipleKeys(t *testing.T) {
	keys := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}
	s := scheduler.NewScheduler(scheduler.SchedulerOption{Debug: true, LogFunc: t.Log})
	for _, key := range keys {
		s.StartWithReturn(key, time.Duration(1*time.Second), func() {})
	}
	for _, key := range keys {
		if s.StartWithReturn(key, time.Duration(1*time.Second), func() {}) {
			t.Error("TestStartHandledCorrectlyWithMultipleKeys failed: Start should return false")
			return
		}
	}
	time.Sleep(2 * time.Second)
	for _, key := range keys {
		if !s.StartWithReturn(key, time.Duration(1*time.Second), func() {}) {
			t.Error("TestStartHandledCorrectlyWithMultipleKeys failed: Start should return true")
			return
		}
	}
	t.Log("TestStartHandledCorrectlyWithMultipleKeys passed")
}

func TestStartHandledTaskCorrectly(t *testing.T) {
	type Integer struct {
		value int
	}
	key := "1"
	s := scheduler.NewScheduler(scheduler.SchedulerOption{Debug: true, LogFunc: t.Log})
	i := Integer{value: 0}
	increase := func() {
		i.value++
	}
	for i := 0; i < 10; i++ {
		s.Start(key, time.Duration(1*time.Second), increase)
	}
	time.Sleep(2 * time.Second)
	if i.value != 1 {
		t.Error("TestStartHandledTaskCorrectly failed: i.value should be 1")
		return
	}
	t.Log("TestStartHandledTaskCorrectly passed")
}

func TestStartHandledTaskCorrectly1(t *testing.T) {
	key := "1"
	s := scheduler.NewScheduler(scheduler.SchedulerOption{Debug: true, LogFunc: t.Log})
	count := 0
	c := make(chan bool, 100)
	for i := 0; i < 100; i++ {
		s.Start(key, time.Duration(1*time.Second), func() {}, c)
	}
	for i := 0; i < 100; i++ {
		if <-c {
			count++
		}
	}
	time.Sleep(2 * time.Second)
	if count != 1 {
		t.Error("TestStartHandledTaskCorrectly failed: count should be 1")
		return
	}
	t.Log("TestStartHandledTaskCorrectly passed")
}

func TestStartHandledTaskCorrectly2(t *testing.T) {
	type Integer struct {
		value int
	}
	key := "1"
	s := scheduler.NewScheduler(scheduler.SchedulerOption{Debug: true, LogFunc: t.Log})
	i := Integer{value: 0}
	store1 := func() {
		i.value = 1
	}
	store2 := func() {
		i.value = 2
	}
	s.StartWithReturn(key, time.Duration(1*time.Second), store1)
	s.StartWithReturn(key, time.Duration(1*time.Second), store2)
	time.Sleep(2 * time.Second)
	if i.value == 1 {
		t.Log("TestStartHandledTaskCorrectly2 passed: i.value =", i.value)
	} else {
		t.Error("TestStartHandledTaskCorrectly2 failed: i.value should be 1")
	}
}

func TestStartHandledTaskCorrectly3(t *testing.T) {
	type Integer struct {
		value int
	}
	key := "1"
	s := scheduler.NewScheduler(scheduler.SchedulerOption{Debug: true, LogFunc: t.Log})
	i := Integer{value: 0}
	store1 := func() {
		i.value = 1
	}
	store2 := func() {
		i.value = 2
	}
	s.StartWithReturn(key, time.Duration(1*time.Second), store1)
	s.ForceStartTaskWithReturn(key, time.Duration(1*time.Second), store2)
	time.Sleep(2 * time.Second)
	if i.value == 2 {
		t.Log("TestStartHandledTaskCorrectly3 passed: i.value =", i.value)
	} else {
		t.Error("TestStartHandledTaskCorrectly3 failed: i.value should be 2")
	}
}

func TestStartHandledTaskCorrectly4(t *testing.T) {
	type Integer struct {
		value int
	}
	key := "1"
	s := scheduler.NewScheduler(scheduler.SchedulerOption{Debug: true, LogFunc: t.Log})
	i := Integer{value: 0}
	store1 := func() {
		i.value = 1
	}
	store2 := func() {
		i.value = 2
	}
	for i := 0; i < 10; i++ {
		s.ForceStartTask(key, time.Duration(1*time.Second), store1)
	}
	time.Sleep(500 * time.Millisecond)
	s.ForceStartTask(key, time.Duration(1*time.Second), store2)
	time.Sleep(2 * time.Second)
	if i.value == 2 {
		t.Log("TestStartHandledTaskCorrectly4 passed: i.value =", i.value)
	} else {
		t.Error("TestStartHandledTaskCorrectly4 failed: i.value should be 2")
	}
}
