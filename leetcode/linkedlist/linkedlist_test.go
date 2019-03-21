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
	cache := LRUConstructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	n := cache.Get(1)
	fmt.Println(n)
	cache.Put(3, 3)
	n = cache.Get(2)
	fmt.Println(n)
	cache.Put(4, 4)
	n = cache.Get(1)
	fmt.Println(n)
	n = cache.Get(3)
	fmt.Println(n)
	n = cache.Get(4)
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

func TestLFUCache(t *testing.T) {

}

func PrintList(head *ListNode) {

	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func TestMyLinkedList(t *testing.T) {
	linkedList := Constructor()
	//linkedList.AddAtHead(1)
	//linkedList.AddAtTail(3)
	////fmt.Println(linkedList.head.Val, linkedList.head.Next)
	//PrintList(linkedList.head)
	//
	//linkedList.AddAtIndex(1, 2)  // linked list becomes 1->2->3
	//PrintList(linkedList.head)
	//
	//linkedList.Get(1)            // returns 2
	//linkedList.DeleteAtIndex(1)  // now the linked list is 1->3
	//PrintList(linkedList.head)
	//
	//n := linkedList.Get(1)            // returns 3
	//fmt.Println(n)
	linkedList.Get(0)
	linkedList.AddAtIndex(1, 2)
	linkedList.Get(0)
	linkedList.Get(1)
	PrintList(linkedList.head)
	linkedList.AddAtIndex(0, 1)
	linkedList.Get(0)
	linkedList.Get(1)
	PrintList(linkedList.head)
}
