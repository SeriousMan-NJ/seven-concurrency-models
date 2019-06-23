package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

func numbers() <-chan int {
	ch := make(chan int)

	go func() {
		i := 2
		for {
			ch <- i
			i++
		}
	}()

	return ch
}

func getPrimes() <-chan int {
	primes := make(chan int)

	go func() {
		ch := numbers()
		for {
			prime, ok := <-ch

			if !ok {
				close(primes)
				return
			}

			primes <- prime
			tmp := make(chan int)

			go func(ch <-chan int, prime int, tmp chan int) {
				for {
					n := <-ch

					if n%prime != 0 {
						tmp <- n
					}
				}
				// close(tmp)
			}(ch, prime, tmp)
			ch = tmp
		}
	}()

	return primes
}

func limit(sec int) <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		time.Sleep(time.Duration(sec) * time.Second)
		ch <- true
	}()

	return ch
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if len(os.Args) < 2 {
		panic("argument must be provided")
	}

	sec, err := strconv.Atoi(os.Args[1])

	if err != nil {
		panic("argument must be integer")
	}

	primes := getPrimes()
	limit := limit(sec)

	for {
		select {
		case <-limit:
			return
		case p := <-primes:
			fmt.Println(p)
		}
	}
}
