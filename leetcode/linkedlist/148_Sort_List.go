package linkedlist

// 148.Sort List
// Sort a linked list in O(n log n) time using constant space complexity.
//
// Example 1:
// Input: 4->2->1->3
// Output: 1->2->3->4
//
// Example 2:
// Input: -1->5->3->4->0
// Output: -1->0->3->4->5
func SortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 快慢指针
	slow := head
	fast := head.Next
	// 找到链表的中点mid，将链表分成两部分
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	mid := slow.Next
	slow.Next = nil

	return merge(SortList(head), SortList(mid))
}

func merge(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	ptr := head
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			ptr.Next = l2
			l2 = l2.Next
		} else {
			ptr.Next = l1
			l1 = l1.Next
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

// 非递归
func SortList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	length := 1
	cur := head
	for cur.Next != nil { // 求链表长度
		length++
		cur = cur.Next
	}

	start := &ListNode{}
	start.Next = head
	var left, right, tail *ListNode
	for n := 1; n < length; n <<= 1 {
		cur = start.Next
		tail = start
		for cur != nil {
			left = cur
			right = split(left, n)
			cur = split(right, n)
			begin, end := merge2(left, right)

			tail.Next = begin
			tail = end
		}
	}

	return start.Next
}

// 将链表分为两部分，前n个和剩余部分，返回剩余部分的头指针
func split(head *ListNode, n int) *ListNode {
	for head != nil && n > 0 { // 向后走n个节点
		head = head.Next
		n--
	}
	var rest *ListNode
	if head != nil {
		rest = head.Next
		head.Next = nil
	}

	return rest
}

// 合并两个链表，返回合并后的链表的头指针和尾指针
func merge2(l1, l2 *ListNode) (*ListNode, *ListNode) {
	head := &ListNode{}
	tail := head

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			tail.Next = l2
			l2 = l2.Next
		} else {
			tail.Next = l1
			l1 = l1.Next
		}
		tail = tail.Next
	}
	if l1 != nil {
		tail.Next = l1
	}
	if l2 != nil {
		tail.Next = l2
	}

	return head.Next, tail
}
