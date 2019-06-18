package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
}

func (c *Counter) Increment() {
	c.count++
}

func (c *Counter) GetCount() int {
	return c.count
}

type Counter2 struct {
	count int
	lock  sync.Mutex
}

func (c *Counter2) Increment() {
	c.lock.Lock()
	c.count++
	c.lock.Unlock()
}

func (c *Counter2) GetCount() int {
	c.lock.Lock()
	tmp := c.count
	c.lock.Unlock()

	return tmp
}

func main() {
	counter := new(Counter)
	counter2 := new(Counter2)
	wg := new(sync.WaitGroup)

	for i := 0; i < 20000; i++ {
		wg.Add(2)
		go func() {
			counter.Increment()
			wg.Done()
		}()
		go func() {
			counter2.Increment()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter.GetCount())
	fmt.Println(counter2.GetCount())
}
