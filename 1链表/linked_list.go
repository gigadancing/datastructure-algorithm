package main

import "fmt"

//
type Node struct {
	val  int
	next *Node
}

//
func NewNode(v int) *Node {
	return &Node{
		val: v,
	}
}

// 打印链表
func PrintList(head *Node) {
	for head != nil {
		fmt.Printf("%v ", head.val)
		head = head.next
	}
	fmt.Println()
}

// 求链长度
func getListLen(head *Node) int {
	var length int
	for head != nil {
		length++
		head = head.next
	}
	return length
}

// 移动较长的链的头
func forwardLongList(head *Node, delta int) *Node {
	for i := 0; i < delta; i++ {
		head = head.next
	}
	return head
}

// 1.链表逆序
// 输入链表的头节点，返回逆序后的头节点
// 思路：按原顺序遍历链表，遍历的同时将节点逆序
func ReverseList(head *Node) *Node {
	if head == nil {
		return nil
	}
	var newHead *Node
	for head != nil {
		// 备份当前节点后面的节点
		next := head.next
		// 当前节点的next指向newHead
		head.next = newHead
		// newHead后移
		newHead = head
		// 当前节点后移
		head = next
	}
	return newHead
}

// 2.求链表逆序指定区间
// 给出链表头节点，逆序指定的区间内的节点
// 1 <= m <= n <= 链表长度
func ReverseListBySection(head *Node, m, n int) *Node {
	if head == nil {
		return nil
	}
	// 需要逆序节点的个数
	changeLen := n - m + 1
	var preHead *Node = nil // m的前驱节点
	var result = head       // 最终转换后的头节点

	// 将preHead移动至m-1位置
	for i := 1; i < m; i++ {
		preHead = head
		head = head.next
	}

	modifyListTail := head // m逆序后的子链表的最后一个节点
	var newHead *Node = nil
	// 逆序
	for ; head != nil && changeLen > 0; changeLen-- {
		next := head.next
		head.next = newHead
		newHead = head
		head = next
	}

	modifyListTail.next = head
	if preHead != nil {
		preHead.next = newHead
	} else {
		result = newHead
	}

	return result
}

// 3.求两链表的交点
// 已知链表A的头节点headA，链表B的头节点headB，两链表相交，求交点对应的节点
func GetIntersectionNode(headA, headB *Node) *Node {
	if headA == nil || headB == nil {
		return nil
	}
	//
	lenA := getListLen(headA)
	lenB := getListLen(headB)
	//
	if lenA > lenB {
		headA = forwardLongList(headA, lenA-lenB)
	} else {
		headB = forwardLongList(headB, lenB-lenA)
	}
	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.next
		headB = headB.next
	}
	return nil
}

// 4.求环起始节点
// 已知链表可能有环，若有环返回环的起始节点，否则返回nil
func DetectCycle(head *Node) *Node {
	if head == nil {
		return nil
	}
	// 快慢指针
	var slow, fast = head, head
	// 相遇节点
	var meet *Node = nil
	for fast != nil {
		// 先个走一步
		slow = slow.next
		fast = fast.next
		// 走到链表末尾
		if fast == nil {
			return nil
		}
		fast = fast.next
		// fast与slow相遇
		if fast == slow {
			meet = fast
			break
		}
	}
	// 没有相遇，证明没有环
	if meet == nil {
		return nil
	}
	for head != nil && meet != nil {
		// 当head和meet相遇，说明到环的起始位置
		if head == meet {
			return head
		}
		// 后移
		head = head.next
		meet = meet.next
	}

	return nil
}

// 5.链表划分
// 已知链表头指针head与数值x，将所有小于x的节点放在大于或等于x的节点前，且保持原来的相对位置
func Partition(head *Node, x int) *Node {
	if head == nil {
		return nil
	}
	// 临时头节点
	var lessHead, greaterHead = NewNode(0), NewNode(0)
	var lessPtr, greaterPtr = lessHead, greaterHead
	for head != nil {
		next := head.next // 备份后面的节点
		if head.val < x {
			lessPtr.next = head
			lessPtr = lessPtr.next
		} else {
			greaterPtr.next = head
			greaterPtr = greaterPtr.next
		}
		head = next
	}
	// 将小于x的节点和大于等于x的节点拼接起来
	lessPtr.next = greaterHead.next
	// 将大于x的最后一个节点的next置为nil，否则会形成环
	greaterPtr.next = nil
	return lessHead.next
}

//
type RandomListNode struct {
	val        int
	next, rand *RandomListNode
}

//
func NewRandomListNode(v int) *RandomListNode {
	return &RandomListNode{
		val: v,
	}
}

// 6.链表深拷贝
// 已知一个复杂的链表，节点中有一个指向本链任意某个节点的随机指针（可以为nil），求该链的深拷贝
func CopyRandomList(head *RandomListNode) *RandomListNode {
	if head == nil {
		return nil
	}
	// 原链表地址和节点位置的映射
	nodeMap := make(map[*RandomListNode]int)
	// 新链每个节点的地址
	nodes := make([]*RandomListNode, 0)
	// 遍历原链表，构建节点地址与节点位置的映射关系
	ptr := head
	for i := 0; ptr != nil; i++ {
		// 创建新节点
		nodes = append(nodes, NewRandomListNode(ptr.val))
		// 节点地址和位置的映射
		nodeMap[ptr] = i
		ptr = ptr.next
	}

	// 最后追加一个nil，防止nodes[i+1]
	nodes = append(nodes, nil)

	// 再次遍历原链表，连接next和rand指针
	ptr = head
	for i := 0; ptr != nil; i++ {
		// 连接next指针
		nodes[i].next = nodes[i+1]
		// 连接rand指针
		if ptr.rand != nil {
			// 原链表中节点rand指针指向节点的位置
			id := nodeMap[ptr.rand]
			// 新节点的rand指向原节点对应的位置
			nodes[i].rand = nodes[id]
		}
		ptr = ptr.next
	}
	return nodes[0]
}

// 7.合并链表
// 已知两个已排序链表的头节点指针h1和h2，将这两个链表合并，合并后的链表仍有序，返回合并后的头节点指针
func MergeTwoLists(h1, h2 *Node) *Node {
	if h1 == nil && h2 != nil {
		return h2
	}
	if h1 != nil && h2 == nil {
		return h1
	}
	if h1 == nil && h2 == nil {
		return nil
	}
	// 临时头节点
	tmpHead := NewNode(-1)
	// ptr指向临时头节点
	ptr := tmpHead
	for h1 != nil && h2 != nil {
		if h1.val < h2.val {
			ptr.next = h1
			h1 = h1.next
		} else {
			ptr.next = h2
			h2 = h2.next
		}
		ptr = ptr.next
	}
	// h1有剩余
	if h1 != nil {
		ptr.next = h1
	}
	// h2有剩余
	if h2 != nil {
		ptr.next = h2
	}

	return tmpHead.next
}
