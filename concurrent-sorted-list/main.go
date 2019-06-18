package main

import (
	"fmt"
	"sync"
)

type Node struct {
	value int
	prev  *Node
	next  *Node
	lock  sync.Mutex
}

func NewNode(value int, prev, next *Node) *Node {
	n := new(Node)
	n.value = value
	n.prev = prev
	n.next = next
	return n
}

type ConcurrentSortedList struct {
	head, tail *Node
}

func NewConcurrentSortedList() *ConcurrentSortedList {
	l := new(ConcurrentSortedList)
	h := new(Node)
	t := new(Node)
	l.head = h
	l.tail = t
	h.next = t
	t.prev = h
	return l
}

func (l *ConcurrentSortedList) Insert(value int) {
	current := l.head
	current.lock.Lock()
	next := current.next
	for {
		next.lock.Lock()
		if next == l.tail || next.value < value {
			node := NewNode(value, current, next)
			next.prev = node
			current.next = node
			current.lock.Unlock()
			next.lock.Unlock()
			return
		}
		current.lock.Unlock()
		current = next
		next = current.next
	}
}

func (l *ConcurrentSortedList) Size() int {
	current := l.tail
	count := 0

	for current.prev != l.head {
		lock := &(current.lock)
		lock.Lock()
		count++
		current = current.prev
		lock.Unlock()
	}

	return count
}

func (l *ConcurrentSortedList) Print() {
	current := l.head.next
	for {
		if current == l.tail {
			return
		} else {
			lock := &(current.lock)
			lock.Lock()
			fmt.Println(current.value)
			current = current.next
			lock.Unlock()
		}
	}
}

func main() {
	l := NewConcurrentSortedList()
	wg := new(sync.WaitGroup)

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func(i int) {
			l.Insert(i)
			wg.Done()
		}(i)
	}

	wg.Wait()

	l.Print()
	fmt.Println("Size:", l.Size())
}
