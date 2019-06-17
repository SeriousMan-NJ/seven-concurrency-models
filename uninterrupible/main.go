package main

import (
	"sync"
	"time"
)

type Object struct {
	lock sync.Mutex
}

func main() {
	o1, o2 := new(Object), new(Object)

	wg := new(sync.WaitGroup)

	wg.Add(2)
	go func() {
		o1.lock.Lock()
		time.Sleep(time.Second)
		o2.lock.Lock()
		o2.lock.Unlock()
		o1.lock.Unlock()
		wg.Done()
	}()

	go func() {
		o2.lock.Lock()
		time.Sleep(time.Second)
		o1.lock.Lock()
		o1.lock.Unlock()
		o2.lock.Unlock()
		wg.Done()
	}()

	wg.Wait()
}
