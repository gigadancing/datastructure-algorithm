package linkedlist

// 460.LFU Cache
// Design and implement a data structure for Least Frequently Used (LFU) cache. It should support the following
// operations: get and put.
//
// get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
// put(key, value) - Set or insert the value if the key is not already present. When the cache reaches its capacity, it
// should invalidate the least frequently used item before inserting a new item. For the purpose of this problem, when
// there is a tie (i.e., two or more keys that have the same frequency), the least recently used key would be evicted.
//
// Follow up:
// Could you do both operations in O(1) time complexity?
//
// Example:
// LFUCache cache = new LFUCache( 2 /* capacity */ );
//
// cache.put(1, 1);
// cache.put(2, 2);
// cache.get(1);       // returns 1
// cache.put(3, 3);    // evicts key 2
// cache.get(2);       // returns -1 (not found)
// cache.get(3);       // returns 3.
// cache.put(4, 4);    // evicts key 1.
// cache.get(1);       // returns -1 (not found)
// cache.get(3);       // returns 3
// cache.get(4);       // returns 4
//
// Least Frequently Used (LFU)，最不经常使用，即淘汰使用次数最少的
//

type LFUNodeList struct {
	head *LFUNode
}

type LFUNode struct {
	key  int      // 键
	val  int      // 值
	freq int      // 访问次数
	prev *LFUNode // 前驱节点
	next *LFUNode // 后继节点
}

type LFUCache struct {
	keyToNode  map[int]*LFUNode     // key -> node
	freqToList map[int]*LFUNodeList // freq -> nodes with the freq
	cap        int
	minFreq    int
}

func LFUConstructor(capacity int) LFUCache {
	return LFUCache{
		keyToNode:  make(map[int]*LFUNode, capacity),
		freqToList: make(map[int]*LFUNodeList, 0),
		cap:        capacity,
	}
}

func (lfu *LFUCache) Get(key int) int {
	return -1
}

func (lfu *LFUCache) Put(key int, value int) {

}

func (lfu *LFUCache) remove(node *LFUNode) {

}

func (lfu *LFUCache) add(node *LFUNode) {

}
