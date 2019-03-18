package linkedlist

// 21.Merge Two Sorted Lists
// Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes
// of the first two lists.
// Example:
//
// Input: 1->2->4, 1->3->4
// Output: 1->1->2->3->4->4

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil && l2 != nil {
		return l2
	}
	if l1 != nil && l2 == nil {
		return l1
	}

	head := &ListNode{}
	ptr := head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			ptr.Next = l1
			l1 = l1.Next
		} else {
			ptr.Next = l2
			l2 = l2.Next
		}
		ptr = ptr.Next
	}

	if l1 != nil {
		ptr.Next = l1
	}
	if l2 != nil {
		ptr.Next = l2
	}

	return head.Next
}

// 递归
func MergeTwoLists2(l1, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil && l2 != nil {
		return l2
	}
	if l1 != nil && l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = MergeTwoLists2(l1.Next, l2)
		return l1
	} else {
		l2.Next = MergeTwoLists2(l1, l2.Next)
		return l2
	}
}

// 141. Linked List Cycle
// Given a linked list, determine if it has a cycle in it.
// To represent a cycle in the given linked list, we use an integer pos which represents the position (0-indexed) in the
// linked list where tail connects to. If pos is -1, then there is no cycle in the linked list.
//
// Example 1:
// Input: head = [3,2,0,-4], pos = 1
// Output: true
//
// Explanation: There is a cycle in the linked list, where tail connects to the second node.
// Example 2:
// Input: head = [1,2], pos = 0
// Output: true
// Explanation: There is a cycle in the linked list, where tail connects to the first node.
//
// Example 3:
// Input: head = [1], pos = -1
// Output: false
// Explanation: There is no cycle in the linked list.
// 用一个集合标记已经遍历过的节点
func HasCycle(head *ListNode) bool {
	mp := make(map[*ListNode]int, 0)
	for head != nil {
		if _, ok := mp[head]; ok {
			return true
		}
		mp[head] = 0
		head = head.Next
	}
	return false
}

// 快慢指针
func HasCycle2(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast, slow := head, head
	for fast != nil {
		if fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
