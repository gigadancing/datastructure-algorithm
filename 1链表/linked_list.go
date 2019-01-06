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
	var front, behind *Node
	tmp := head
	for i := 1; i < n; i++ {
		if i == m-1 {
			front = tmp // m-1位置的节点
		}
		tmp = tmp.next // n位置的节点
	}
	behind = tmp.next // 备份n以后的节点
	tmp.next = nil
	subListHead := ReverseList(front.next) // 逆序[m, n]区间的节点
	// 找到[m, n]最后一个节点tail
	tail := subListHead
	for tail.next != nil {
		tail = tail.next
	}
	// 连接
	tail.next = behind
	front.next = subListHead
	return head
}

func main() {
	node1 := NewNode(1)
	node2 := NewNode(2)
	node3 := NewNode(3)
	node4 := NewNode(4)
	node5 := NewNode(5)
	node6 := NewNode(6)
	node7 := NewNode(7)
	node8 := NewNode(8)
	node9 := NewNode(9)
	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5
	node5.next = node6
	node6.next = node7
	node7.next = node8
	node8.next = node9

	fmt.Println("逆序前：")
	PrintList(node1)

	fmt.Println("逆序后：")
	newHead := ReverseList(node1)
	PrintList(newHead)

	fmt.Println("逆序指定区间：")
	newHead = ReverseListBySection(newHead, 2, 4)
	PrintList(newHead)
}
