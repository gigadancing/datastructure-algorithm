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

// 哈希表的表长一般为质数，使取余的结果非常分散，不易重复

// 例1. 最长回文串
// 已知一个只包含大小写字符的字符串，求用该字符串中的字符可以生成的最长回文字符串长度。
// 例如：s="abccccddaa"，可以生成的最长回文串长度为9，如"dccaaaccd"、"adccbccda"、"acdcacdca"等都是正确的。
func LongestPalindrome(s string) int {
	bytes := [128]byte{0}
	maxLength := 0 // 回文字符串中偶数部分
	flag := 0      // 是否有中心点
	for _, v := range s {
		bytes[v-'0']++
	}
	for i := 0; i < 128; i++ {
		if bytes[i]%2 == 0 {
			maxLength += int(bytes[i])
		} else {
			maxLength += int(bytes[i]) - 1
			flag = 1
		}
	}

	return maxLength + flag
}
