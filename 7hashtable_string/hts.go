package hts

// 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// 链表节点构造函数
func NewListNode(v int) *ListNode {
	return &ListNode{
		Val: v,
	}
}

// 哈希函数
func HashFunc(key, tableLen int) int {
	return key % tableLen
}

// 插入节点
func Insert(hashTable []*ListNode, node *ListNode, tableLen int) {
	index := HashFunc(node.Val, tableLen)
	node.Next = hashTable[index]
	hashTable[index] = node
}

// 查找值
func Search(hashTable []*ListNode, tableLen, value int) bool {
	index := HashFunc(value, tableLen)
	if hashTable[index] == nil {
		return false
	}
	for p := hashTable[index]; p != nil; p = p.Next {
		if p.Val == value {
			return true
		}
	}
	return false
}
