package main

import (
	"fmt"
	"testing"
)

//
func TestReverseList(t *testing.T) {
	node1 := NewNode(1)
	node2 := NewNode(2)
	node3 := NewNode(3)
	node4 := NewNode(4)
	node5 := NewNode(5)
	node6 := NewNode(6)
	node7 := NewNode(7)
	node8 := NewNode(8)
	node9 := NewNode(9)
	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5
	node5.next = node6
	node6.next = node7
	node7.next = node8
	node8.next = node9
	fmt.Println("逆序前：")
	PrintList(node1)
	fmt.Println("逆序后：")
	newHead := ReverseList(node1)
	PrintList(newHead)
}

//
func TestReverseListBySection(t *testing.T) {
	node1 := NewNode(1)
	node2 := NewNode(2)
	node3 := NewNode(3)
	node4 := NewNode(4)
	node5 := NewNode(5)
	node6 := NewNode(6)
	node7 := NewNode(7)
	node8 := NewNode(8)
	node9 := NewNode(9)
	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5
	node5.next = node6
	node6.next = node7
	node7.next = node8
	node8.next = node9
	fmt.Println("逆序指定区间：")
	newHead := ReverseListBySection(node1, 1, 4)
	PrintList(newHead)
}

//
func TestGetIntersectionNode(t *testing.T) {

}

//
func TestDetectCycle(t *testing.T) {
	node1 := NewNode(1)
	node2 := NewNode(2)
	node3 := NewNode(3)
	node4 := NewNode(4)
	node5 := NewNode(5)
	node6 := NewNode(6)
	node7 := NewNode(7)
	node8 := NewNode(8)
	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5
	node5.next = node6
	node6.next = node7
	node7.next = node8
	node8.next = node6
	n := DetectCycle(node1)
	fmt.Println(n.val)
}
