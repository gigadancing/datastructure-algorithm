package linkedlist

// 707.Design Linked List
// Design your implementation of the linked list. You can choose to use the singly linked list or the doubly linked list.
// A node in a singly linked list should have two attributes: val and next. val is the value of the current node, and
// next is a pointer/reference to the next node. If you want to use the doubly linked list, you will need one more
// attribute prev to indicate the previous node in the linked list. Assume all nodes in the linked list are 0-indexed.
//
// Implement these functions in your linked list class:
//
// get(index) : Get the value of the index-th node in the linked list. If the index is invalid, return -1.
// addAtHead(val) : Add a node of value val before the first element of the linked list. After the insertion, the new
//                  node will be the first node of the linked list.
// addAtTail(val) : Append a node of value val to the last element of the linked list.
// addAtIndex(index, val) : Add a node of value val before the index-th node in the linked list. If index equals to the
//                          length of linked list, the node will be appended to the end of linked list. If index is
//                          greater than the length, the node will not be inserted.
// deleteAtIndex(index) : Delete the index-th node in the linked list, if the index is valid.
//
// Example:
//
// MyLinkedList linkedList = new MyLinkedList();
// linkedList.addAtHead(1);
// linkedList.addAtTail(3);
// linkedList.addAtIndex(1, 2);  // linked list becomes 1->2->3
// linkedList.get(1);            // returns 2
// linkedList.deleteAtIndex(1);  // now the linked list is 1->3
// linkedList.get(1);            // returns 3
//
// Note:
// All values will be in the range of [1, 1000].
// The number of operations will be in the range of [1, 1000].
// Please do not use the built-in LinkedList library.
type MyLinkedList struct {
	head *ListNode
	tail *ListNode
	size int
}

// Initialize your data structure here.
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

// Get the value of the index-th node in the linked list. If the index is invalid, return -1.
func (this *MyLinkedList) Get(index int) int {
	if index >= this.size || index < 0 || index > 1000 {
		return -1
	}
	p := this.head
	for index > 0 && p != nil {
		p = p.Next
		index--
	}
	if p != nil {
		return p.Val
	}
	return -1
}

// Add a node of value val before the first element of the linked list. After the insertion, the new node will be the
// first node of the linked list.
func (this *MyLinkedList) AddAtHead(val int) {
	if val < 0 || val > 1000 {
		return
	}
	node := &ListNode{Val: val}
	if this.size == 0 {
		this.head = node
		this.tail = node
	} else {
		node.Next = this.head
		this.head = node
	}
	this.size++
}

// Append a node of value val to the last element of the linked list.
func (this *MyLinkedList) AddAtTail(val int) {
	if val < 0 || val > 1000 {
		return
	}
	node := &ListNode{Val: val}
	if this.size == 0 {
		this.tail = node
		this.head = node
	} else {
		this.tail.Next = node
		this.tail = this.tail.Next
	}
	this.size++
}

// Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list,
// the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted.
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size || index < 0 || val < 0 || val > 1000 {
		return
	}
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	if index == this.size {
		this.AddAtTail(val)
		return
	}

	index -= 1
	p := this.head
	node := &ListNode{Val: val}
	for index > 0 && p != nil {
		index--
		p = p.Next
	}
	next := p.Next
	p.Next = node
	node.Next = next
	this.size++
}

// Delete the index-th node in the linked list, if the index is valid.
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.size || index < 0 {
		return
	}
	if index == 0 {
		next := this.head.Next
		this.head.Next = nil
		this.head = next
	} else {
		p := this.head
		index -= 1
		for p != nil && index > 0 {
			index--
			p = p.Next
		}
		if p.Next == this.tail {
			p.Next = nil
			this.tail = p
		} else {
			p2 := p.Next
			p3 := p.Next.Next
			p.Next = p3
			p2.Next = nil
		}
	}
	this.size--
}
