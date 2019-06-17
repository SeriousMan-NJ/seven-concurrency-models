package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Philosopher struct {
	left, right *Chopstick
	random      *rand.Rand
}

func NewPhilosopher(left, right *Chopstick, random *rand.Rand) *Philosopher {
	p := new(Philosopher)
	p.left = left
	p.right = right
	p.random = random
	return p
}

func (p *Philosopher) Run() {
	for {
		time.Sleep(time.Duration(p.random.Intn(1000)) * time.Millisecond)
		p.left.lock.Lock()
		fmt.Printf("(LEFT) Chopstick #%d is in use\n", p.left.id)
		time.Sleep(time.Second) // 강제로 데드락을 만드는 코드
		p.right.lock.Lock()
		fmt.Printf("(RIGHT) Chopstick #%d is in use\n", p.right.id)
		time.Sleep(time.Duration(p.random.Intn(1000)) * time.Millisecond)
		p.right.lock.Unlock()
		fmt.Printf("(RIGHT) Chopstick #%d is free\n", p.right.id)
		p.left.lock.Unlock()
		fmt.Printf("(LEFT) Chopstick #%d is in use\n", p.left.id)
	}
}
