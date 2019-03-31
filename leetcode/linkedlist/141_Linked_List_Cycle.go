package linkedlist

// 141. Linked List Cycle I
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
