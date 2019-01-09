package main

import (
	"fmt"
	"testing"
)

func TestMyStack(t *testing.T) {
	s := NewMyStack()
	s.push(1)
	s.push(2)
	s.push(3)
	s.push(4)
	s.push(5)
	for !s.empty() {
		fmt.Println(s.pop())
	}
}

func TestNewMyQueue(t *testing.T) {
	q := NewMyQueue()
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	q.enqueue(4)
	q.enqueue(5)
	for !q.empty() {
		fmt.Printf("%v ", q.dequeue())
	}
	fmt.Println()
}
