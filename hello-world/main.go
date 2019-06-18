package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		fmt.Println("Hello from go routine")
		wg.Done()
	}()

	fmt.Println("Hello from main routine")
	wg.Wait()
}
