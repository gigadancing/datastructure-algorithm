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
// Least Recently Used (LRU)，最近最少使用，即淘汰最近不使用的
type LRUNode struct {
	Key, Val   int
	Prev, Next *LRUNode
}

type LRUCache struct {
	Store      map[int]*LRUNode
	Head, Tail *LRUNode // 双向链表的头尾节点
	Cap        int      // cache的容量
}

func LRUConstructor(capacity int) LRUCache {
	return LRUCache{
		Store: make(map[int]*LRUNode, capacity),
		Cap:   capacity,
	}
}

func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.Store[key]; ok { // key存在
		lru.Remove(node)
		lru.AddFront(node)
		return node.Val
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	node, ok := lru.Store[key]
	if ok {
		node.Val = value
		lru.Remove(node)
		lru.AddFront(node)
	} else {
		node = &LRUNode{Key: key, Val: value}
		lru.Store[key] = node
		lru.AddFront(node)
	}
	if len(lru.Store) > lru.Cap {
		node = lru.Tail
		lru.Remove(node)
		if node != nil {
			delete(lru.Store, node.Key)
		}
	}
}

// 从链表中删除节点
func (lru *LRUCache) Remove(node *LRUNode) {
	if node == lru.Head {
		lru.Head = node.Next
	}
	if node == lru.Tail {
		lru.Tail = node.Prev
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
}

// 将节点插入到最前面
func (lru *LRUCache) AddFront(node *LRUNode) {
	node.Prev = nil
	if lru.Head == nil { // 当前还没有节点
		lru.Head = node
		lru.Tail = node
		return
	}
	node.Next = lru.Head
	lru.Head.Prev = node
	lru.Head = node
}
