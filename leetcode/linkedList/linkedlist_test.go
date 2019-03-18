package linkedList

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
