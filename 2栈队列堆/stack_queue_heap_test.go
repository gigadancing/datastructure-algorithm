package main

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
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

func TestMyQueue(t *testing.T) {
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

func TestMyStackWithMin(t *testing.T) {
	s := NewMyStackWithMin()

	s.push(5)                       // 5
	fmt.Println("min:", s.getMin()) // 5
	s.push(-1)                      // 5 -1
	fmt.Println("min:", s.getMin()) // -1
	s.pop()                         // 5
	fmt.Println("min:", s.getMin()) // 5
	s.push(2)                       // 5 2
	s.push(8)                       // 5 2 8
	fmt.Println("min:", s.getMin()) // 2
	s.pop()
	fmt.Println("min:", s.getMin()) // 2
}

func TestCheckOrder(t *testing.T) {
	q := queue.New()
	q.Enqueue(3)
	q.Enqueue(2)
	q.Enqueue(5)
	q.Enqueue(4)
	q.Enqueue(1)
	fmt.Println("is valid:", CheckOrder(q))

	q2 := queue.New()
	q2.Enqueue(3)
	q2.Enqueue(1)
	q2.Enqueue(2)
	q2.Enqueue(4)
	q2.Enqueue(5)
	fmt.Println("is valid:", CheckOrder(q2))
}
