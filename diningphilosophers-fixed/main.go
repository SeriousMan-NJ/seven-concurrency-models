package main

import (
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg := new(sync.WaitGroup)

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	c := make([]*Chopstick, 5)
	p := make([]*Philosopher, 5)

	for i := 0; i < 5; i++ {
		c[i] = NewChopstick(i)
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			p[i] = NewPhilosopher(c[i], c[(i+1)%5], r)
			p[i].Run()
			wg.Done()
		}(i)
	}

	wg.Wait()
}
