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

// 链表逆序
// 输入链表的头节点，返回逆序后的头节点
// 思路：按原顺序遍历链表，遍历的同时将节点逆序
func ReverseList(head *Node) *Node {
	if head == nil {
		return nil
	}
	// 处理头节点
	newHead := head
	head = head.next
	newHead.next = nil
	// 从二个节点开始
	for head != nil {
		// 在逆序节点的时候要先把后面的节点保存
		tmp := head.next
		// 逆序节点
		head.next = newHead
		// 将新的头节点后移
		newHead = head
		// 将原头节点后移
		head = tmp
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

func main() {
	node1 := NewNode(1)
	node2 := NewNode(2)
	node3 := NewNode(3)
	node4 := NewNode(4)
	node5 := NewNode(5)
	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5
	PrintList(node1)
	newHead := ReverseList(node1)
	PrintList(newHead)
}
