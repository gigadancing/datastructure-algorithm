package linkedlist

// 142.Linked List Cycle II
// Given a linked list, return the node where the cycle begins. If there is no cycle, return null.
//
// To represent a cycle in the given linked list, we use an integer pos which represents the position (0-indexed) in the
// linked list where tail connects to. If pos is -1, then there is no cycle in the linked list.
// Note: Do not modify the linked list.
//
// Example 1:
// Input: head = [3,2,0,-4], pos = 1
// Output: tail connects to node index 1
// Explanation: There is a cycle in the linked list, where tail connects to the second node.
//
// Example 2:
// Input: head = [1,2], pos = 0
// Output: tail connects to node index 0
// Explanation: There is a cycle in the linked list, where tail connects to the first node.
//
// Example 3:
// Input: head = [1], pos = -1
// Output: no cycle
// Explanation: There is no cycle in the linked list.
func DetectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	for fast != nil {
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next // 走两步
		slow = slow.Next      // 走一步
		if fast == slow {     // 相遇
			break
		}
	}

	if fast == nil {
		return nil
	}

	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
