package main

import (
	"fmt"
	"os"
	"strconv"
)

func getPrimes(limit int) chan int {
	primes := make(chan int)
	numbers := make(chan int)

	go func() {
		for i := 2; i <= limit; i++ {
			numbers <- i
		}
		close(numbers)
	}()

	go func() {
		ch := numbers
		for {
			prime, ok := <-ch

			if !ok {
				close(primes)
				return
			}

			primes <- prime
			tmp := make(chan int)

			go func(ch chan int, prime int, tmp chan int) {
				for n := range ch {
					if n%prime != 0 {
						tmp <- n
					}
				}
				close(tmp)
			}(ch, prime, tmp)
			ch = tmp
		}
	}()

	return primes
}

func main() {
	if len(os.Args) < 2 {
		panic("argument must be provided")
	}

	limit, err := strconv.Atoi(os.Args[1])

	if err != nil {
		panic("argument must be integer")
	}

	for prime := range getPrimes(limit) {
		fmt.Println(prime)
	}
}
