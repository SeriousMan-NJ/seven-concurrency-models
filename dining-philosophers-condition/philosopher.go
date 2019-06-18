package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Philosopher struct {
	id        int
	eating    bool
	left      *Philosopher
	right     *Philosopher
	table     *sync.Mutex
	condition *sync.Cond
	random    *rand.Rand
}

func NewPhilosopher(id int, table *sync.Mutex, random *rand.Rand) *Philosopher {
	p := new(Philosopher)
	p.id = id
	p.eating = false
	p.table = table
	p.condition = sync.NewCond(table)
	p.random = random
	return p
}

func (p *Philosopher) SetLeft(left *Philosopher) {
	p.left = left
}

func (p *Philosopher) SetRight(right *Philosopher) {
	p.right = right
}

func (p *Philosopher) Run() {
	for {
		p.think()
		p.eat()
	}
}

func (p *Philosopher) think() {
	p.table.Lock()
	p.eating = false
	p.left.condition.Signal()
	p.right.condition.Signal()
	p.table.Unlock()
	time.Sleep(time.Second)
}

func (p *Philosopher) eat() {
	p.table.Lock()
	for p.left.eating || p.right.eating {
		p.condition.Wait()
	}
	p.eating = true
	fmt.Printf("%d eating\n", p.id)
	p.table.Unlock()
	time.Sleep(time.Second)
}
