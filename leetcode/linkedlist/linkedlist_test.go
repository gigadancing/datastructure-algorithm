package linkedlist

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	n11 := ListNode{Val: 1}
	n12 := ListNode{Val: 2}
	n13 := ListNode{Val: 4}
	n21 := ListNode{Val: 1}
	n22 := ListNode{Val: 3}
	n23 := ListNode{Val: 4}
	n11.Next = &n12
	n12.Next = &n13
	n21.Next = &n22
	n22.Next = &n23
	h := MergeTwoLists(&n11, &n21)
	for h != nil {
		fmt.Printf("%d ", h.Val)
		h = h.Next
	}
	fmt.Println()
}

func TestMergeTwoLists2(t *testing.T) {
	n11 := ListNode{Val: 1}
	n12 := ListNode{Val: 2}
	n13 := ListNode{Val: 4}
	n21 := ListNode{Val: 1}
	n22 := ListNode{Val: 3}
	n23 := ListNode{Val: 4}
	n11.Next = &n12
	n12.Next = &n13
	n21.Next = &n22
	n22.Next = &n23
	h := MergeTwoLists2(&n11, &n21)
	for h != nil {
		fmt.Printf("%d ", h.Val)
		h = h.Next
	}
	fmt.Println()
}

func TestHasCycle(t *testing.T) {
	n1 := ListNode{Val: 3}
	n2 := ListNode{Val: 2}
	n3 := ListNode{Val: 0}
	n4 := ListNode{Val: 4}
	n1.Next = &n2
	n2.Next = &n3
	n3.Next = &n4
	n4.Next = &n2

	head := HasCycle(&n1)
	fmt.Println(head)
}

func TestHasCycle2(t *testing.T) {
	n1 := ListNode{Val: 3}
	n2 := ListNode{Val: 2}
	n3 := ListNode{Val: 0}
	n4 := ListNode{Val: 4}
	n1.Next = &n2
	n2.Next = &n3
	n3.Next = &n4
	n4.Next = &n2

	head := HasCycle2(&n1)
	fmt.Println(head)
}

func TestDetectCycle(t *testing.T) {
	n1 := ListNode{Val: 3}
	n2 := ListNode{Val: 2}
	n3 := ListNode{Val: 0}
	n4 := ListNode{Val: 4}
	n1.Next = &n2
	n2.Next = &n3
	n3.Next = &n4
	n4.Next = &n2
	meet := DetectCycle(&n1)
	if meet != nil {
		fmt.Println(meet.Val)
	} else {
		fmt.Println("no cycle")
	}

	n5 := ListNode{Val: 1}
	n6 := ListNode{Val: 2}
	n5.Next = &n6
	n6.Next = &n5
	meet = DetectCycle(&n5)
	if meet != nil {
		fmt.Println(meet.Val)
	} else {
		fmt.Println("no cycle")
	}

	n7 := ListNode{Val: 1}
	n8 := ListNode{Val:2}
	n7.Next = &n8
	meet = DetectCycle(&n7)
	if meet != nil {
		fmt.Println(meet.Val)
	} else {
		fmt.Println("no cycle")
	}
}
