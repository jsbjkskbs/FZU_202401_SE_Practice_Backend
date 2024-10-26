package generator_test

import (
	"sfw/pkg/utils/generator"
	"sync"
	"testing"
	"time"
)

func TestSnowFlake(t *testing.T) {
	result := sync.Map{}
	g, _ := generator.NewSnowflake(1)
	for i := 0; i < 1000; i++ {
		go func() {
			id := g.Generate()
			result.Store(id, struct{}{})
		}()
	}
	time.Sleep(1 * time.Second)
	len := 0
	result.Range(func(key, value interface{}) bool {
		len++
		t.Logf("%v %v", key, value)
		return true
	})
	if len != 1000 {
		t.Errorf("expected 1000, got %d", len)
	} else {
		t.Logf("success")
	}
}
