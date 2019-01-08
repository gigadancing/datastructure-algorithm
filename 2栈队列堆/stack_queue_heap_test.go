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
