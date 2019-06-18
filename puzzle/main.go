package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	answerReady := false
	answer := 0

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		answer = 42
		answerReady = true
		wg.Done()
	}()

	go func() {
		if answerReady {
			fmt.Println("The meaning of life is:", answer)
		} else {
			fmt.Println("I don't know the answer")
		}
		wg.Done()
	}()

	wg.Wait()
}
