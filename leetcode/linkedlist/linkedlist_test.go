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
	n8 := ListNode{Val: 2}
	n7.Next = &n8
	meet = DetectCycle(&n7)
	if meet != nil {
		fmt.Println(meet.Val)
	} else {
		fmt.Println("no cycle")
	}
}

func TestLRUCache(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	n := cache.Get(1) // returns 1
	fmt.Println(n)
	cache.Put(3, 3)  // evicts key 2
	n = cache.Get(2) // returns -1 (not found)
	fmt.Println(n)
	cache.Put(4, 4)  // evicts key 1
	n = cache.Get(1) // returns -1 (not found)
	fmt.Println(n)
	n = cache.Get(3) // returns 3
	fmt.Println(n)
	n = cache.Get(4) // returns 4
	fmt.Println(n)
}

func TestSortList(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 1}
	n4 := &ListNode{Val: 3}
	n5 := &ListNode{Val: 9}
	n6 := &ListNode{Val: 7}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	head := SortList(n1)
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func TestSortList2(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 1}
	n4 := &ListNode{Val: 3}
	n5 := &ListNode{Val: 9}
	n6 := &ListNode{Val: 7}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	head := SortList2(n1)
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}
