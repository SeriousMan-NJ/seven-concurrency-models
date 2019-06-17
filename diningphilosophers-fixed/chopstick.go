package main

import "sync"

type Chopstick struct {
	id   int
	lock sync.Mutex
}

func NewChopstick(id int) *Chopstick {
	c := new(Chopstick)
	c.id = id
	return c
}
