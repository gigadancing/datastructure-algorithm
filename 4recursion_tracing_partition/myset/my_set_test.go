package myset

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	s := New()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Add(2)
	fmt.Println("myset data:", s.List())
	s.Clear()
	if s.Empty() {
		fmt.Println("after clear, myset is empty")
	}

	s.Add(4)
	s.Add(5)
	s.Add(6)
	fmt.Println("myset data:", s.List())
	if s.Has(5) {
		fmt.Println("5 does exist")
	}
	s.Remove(4)
	s.Remove(5)
	fmt.Println("remove 4,5:", s.List())
}
