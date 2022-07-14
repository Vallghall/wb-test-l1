package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Counter is a concurrency-safe counter
type Counter struct {
	counter int64
}

// Inc increments the value of the Counter concurrency-safely
func (c *Counter) Inc() {
	atomic.AddInt64(&c.counter, 1)
}

// Reset sets the value of the Counter to 0 concurrency-safely
func (c *Counter) Reset() {
	atomic.StoreInt64(&c.counter, 0)
}

func (c Counter) Value() int64 {
	return atomic.LoadInt64(&c.counter)
}

func main() {
	c := Counter{}
	wg := sync.WaitGroup{}
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}

	wg.Wait()
	fmt.Println(c.Value())
}
