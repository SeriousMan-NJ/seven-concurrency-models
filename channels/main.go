package main

import (
	"fmt"
	"runtime"
)

func ReadAll(ch <-chan int) []int {
	var r []int
	for {
		if v, ok := <-ch; ok {
			r = append(r, v)
		} else {
			return r
		}
	}
}

func WriteAll(ch chan<- int, l []int) {
	for i := 0; i < len(l); i++ {
		ch <- l[i]
	}
	close(ch)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	ch := make(chan int)

	go func() {
		fmt.Println(<-ch)
	}()
	ch <- 3
}
