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

//
func TestPartition(t *testing.T) {
	node1 := NewNode(1)
	node2 := NewNode(3)
	node3 := NewNode(5)
	node4 := NewNode(7)
	node5 := NewNode(9)
	node6 := NewNode(2)
	node7 := NewNode(4)
	node8 := NewNode(6)
	node9 := NewNode(8)
	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5
	node5.next = node6
	node6.next = node7
	node7.next = node8
	node8.next = node9
	h := Partition(node1, 7)
	PrintList(h)
}

//
func TestCopyRandomList(t *testing.T) {
	a := NewRandomListNode(1)
	b := NewRandomListNode(2)
	c := NewRandomListNode(3)
	d := NewRandomListNode(4)
	e := NewRandomListNode(5)

	a.next = b
	b.next = c
	c.next = d
	d.next = e

	a.rand = c
	b.rand = d
	c.rand = c
	e.rand = d

	h := CopyRandomList(a)
	for p := h; p != nil; p = p.next {
		if p.rand != nil {
			fmt.Println("val:", p.val, "--rand-->", p.rand.val)
		}
	}
}

//
func TestMergeTwoLists(t *testing.T) {
	node0 := NewNode(0)
	node1 := NewNode(1)
	node4 := NewNode(4)
	node5 := NewNode(5)
	node6 := NewNode(6)
	node7 := NewNode(7)
	node8 := NewNode(43)
	node9 := NewNode(59)
	node10 := NewNode(32)
	node11 := NewNode(21)
	node12 := NewNode(18)
	// 0->5->7->18
	node0.next = node5
	node5.next = node7
	node7.next = node12
	// 1->4->6->21->32->43->59
	node1.next = node4
	node4.next = node6
	node6.next = node11
	node11.next = node10
	node10.next = node8
	node8.next = node9

	h := MergeTwoLists(node1, node0)
	// 0->1->4->5->6->7>18->21->32->43->59
	PrintList(h)
}
