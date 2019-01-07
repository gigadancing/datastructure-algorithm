package main

import (
	"fmt"
)

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

// 链表逆序
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

// 打印链表
func PrintList(head *Node) {
	for head != nil {
		fmt.Printf("%v ", head.val)
		head = head.next
	}
	fmt.Println()
}

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
func getListLen(head *Node) int {
	var length int
	for head != nil {
		length++
		head = head.next
	}
	return length
}

//
func forwardLongList(head *Node, delta int) *Node {
	for i := 0; i < delta; i++ {
		head = head.next
	}
	return head
}

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
