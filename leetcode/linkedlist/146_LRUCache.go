package linkedlist

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
		lru.AddFront(v)
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
		lru.AddFront(v)
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

func (lru *LRUCache) AddFront(node *LRUNode) {
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
