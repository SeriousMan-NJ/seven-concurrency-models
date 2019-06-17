package main

import (
	"fmt"
	"sync"
)

func HelloWorld() {
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		fmt.Println("Hello from go routine")
		wg.Done()
	}()

	fmt.Println("Hello from main routine")
	wg.Wait()
}
