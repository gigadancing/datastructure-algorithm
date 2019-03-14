package biq

import (
	"fmt"
	"testing"
)

func TestFindAllEquivalentPoints(t *testing.T) {
	nums := []Point{{1, 1}, {1, 4}, {3, 4}, {3, 7}, {9, 7}, {9, 1}}
	res := FindAllEquivalentPoints(nums, 4)
	for _, p := range res {
		fmt.Println("(", p.x, ",", p.y, ")")
	}
}

func TestReverse(t *testing.T) {
	n1 := NewListNode(1)
	n2 := NewListNode(2)
	n3 := NewListNode(3)
	n4 := NewListNode(4)
	n1.next = n2
	n2.next = n3
	n3.next = n4

	head := Reverse(n1)
	PrintList(head)
}

func TestAddList(t *testing.T) {
	n1 := NewListNode(7)
	n2 := NewListNode(2)
	n3 := NewListNode(3)
	n4 := NewListNode(4)
	n1.next = n2
	n2.next = n3
	n3.next = n4

	m1 := NewListNode(9)
	m2 := NewListNode(9)
	m3 := NewListNode(8)
	m4 := NewListNode(7)
	m1.next = m2
	m2.next = m3
	m3.next = m4
	// 7234+9987=17221
	res := AddList(n1, m1)
	PrintList(res)
}
