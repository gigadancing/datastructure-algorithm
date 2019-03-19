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

// 146.LRU Cache
// Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following
// operations: get and put.
//
// get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
// put(key, value) - Set or insert the value if the key is not already present. When the cache reached its capacity, it
// should invalidate the least recently used item before inserting a new item.
// Follow up:
// Could you do both operations in O(1) time complexity?
// Example:
//
// LRUCache cache = new LRUCache( 2 /* capacity */ );
//
// cache.put(1, 1);
// cache.put(2, 2);
// cache.get(1);       // returns 1
// cache.put(3, 3);    // evicts key 2
// cache.get(2);       // returns -1 (not found)
// cache.put(4, 4);    // evicts key 1
// cache.get(1);       // returns -1 (not found)
// cache.get(3);       // returns 3
// cache.get(4);       // returns 4
//
// Your LRUCache object will be instantiated and called as such:
// obj := Constructor(capacity)
// param_1 := obj.Get(key)
// obj.Put(key,value)
//
type LRUNode struct {
	Key, Val   int
	Prev, Next *LRUNode
}

type LRUCache struct {
	Data       map[int]*LRUNode
	Head, Tail *LRUNode // 双向链表的头尾节点
	Cap        int      // cache的容量
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		Data: make(map[int]*LRUNode, capacity),
		Head: &LRUNode{Key: -1, Val: -1, Prev: nil, Next: nil},
		Tail: &LRUNode{Key: -1, Val: -1, Prev: nil, Next: nil},
		Cap:  capacity,
	}
	lru.Head.Next = lru.Tail
	lru.Tail.Prev = lru.Head
	return lru
}

func (lru *LRUCache) Get(key int) int {
	if key <= 0 {
		return -1
	}
	if v, ok := lru.Data[key]; ok {
		lru.LeastRecentlyUsed(v)
		return v.Val
	}

	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	if v, ok := lru.Data[key]; ok { // key已存在
		v.Val = value
		if lru.Head.Next == v { // key就是最近使用的
			return
		}
		lru.LeastRecentlyUsed(v)
		return
	}

	if len(lru.Data) == lru.Cap { // cache已满
		prev := lru.Tail.Prev
		pprev := lru.Tail.Prev.Prev
		pprev.Next = lru.Tail
		lru.Tail.Prev = pprev
		delete(lru.Data, prev.Key) // 删除key-value
		lru.Put(key, value)        // 加入key-value
	} else { // cache未满
		node := &LRUNode{Key: key, Val: value}
		next := lru.Head.Next
		lru.Head.Next = node
		node.Prev = lru.Head
		node.Next = next
		next.Prev = node
		lru.Data[key] = node
	}
}

func (lru *LRUCache) LeastRecentlyUsed(node *LRUNode) {
	prev := node.Prev
	next := node.Next
	prev.Next = next
	next.Prev = prev
	headNext := lru.Head.Next
	lru.Head.Next = node
	node.Prev = lru.Head
	node.Next = headNext
	headNext.Prev = node
}
