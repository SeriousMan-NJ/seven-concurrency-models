package main

import (
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	philosophers := new([5]*Philosopher)
	table := new(sync.Mutex)
	wg := new(sync.WaitGroup)

	for i := 0; i < 5; i++ {
		philosophers[i] = NewPhilosopher(i, table, r)
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		philosophers[i].SetLeft(philosophers[(i+4)%5])
		philosophers[i].SetRight(philosophers[(i+1)%5])
		go func(i int) {
			philosophers[i].Run()
			wg.Done()
		}(i)
	}

	wg.Wait()
}
