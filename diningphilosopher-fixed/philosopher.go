package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Philosopher struct {
	first, second *Chopstick
	random        *rand.Rand
}

func NewPhilosopher(left, right *Chopstick, random *rand.Rand) *Philosopher {
	p := new(Philosopher)
	if left.id < right.id {
		p.first, p.second = left, right
	} else {
		p.second, p.first = left, right
	}
	p.random = random
	return p
}

func (p *Philosopher) Run() {
	for {
		time.Sleep(time.Duration(p.random.Intn(1000)) * time.Millisecond)
		p.first.lock.Lock()
		fmt.Printf("(LEFT) Chopstick #%d is in use\n", p.first.id)
		time.Sleep(time.Second) // 이제 데드락에 걸리지 않는다.
		p.second.lock.Lock()
		fmt.Printf("(RIGHT) Chopstick #%d is in use\n", p.second.id)
		time.Sleep(time.Duration(p.random.Intn(1000)) * time.Millisecond)
		p.second.lock.Unlock()
		fmt.Printf("(RIGHT) Chopstick #%d is free\n", p.second.id)
		p.first.lock.Unlock()
		fmt.Printf("(LEFT) Chopstick #%d is in use\n", p.first.id)
	}
}
