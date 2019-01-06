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
