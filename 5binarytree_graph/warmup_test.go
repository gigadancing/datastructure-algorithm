package bg

import (
	"fmt"
	"testing"
)

func TestNewTreeNode(t *testing.T) {
	root := NewTreeNode(0)
	a := NewTreeNode(1)
	b := NewTreeNode(2)
	c := NewTreeNode(3)
	d := NewTreeNode(4)
	e := NewTreeNode(5)
	f := NewTreeNode(6)
	root.left = a
	root.right = b
	a.left = c
	a.right = d
	b.left = e
	b.right = f
	PreorderPrint(root, 0)
}

func TestTraverse(t *testing.T) {
	root := NewTreeNode(0)
	a := NewTreeNode(1)
	b := NewTreeNode(2)
	c := NewTreeNode(3)
	d := NewTreeNode(4)
	e := NewTreeNode(5)
	f := NewTreeNode(6)
	g := NewTreeNode(7)
	h := NewTreeNode(8)
	i := NewTreeNode(9)
	j := NewTreeNode(10)
	root.left = a
	root.right = b
	a.left = c
	a.right = d
	b.left = e
	b.right = f
	d.left = g
	e.right = h
	f.left = i
	f.right = j
	path := make([]*TreeNode, 0)
	result := make([][]*TreeNode, 0)
	Traverse(root, path, &result)
	for _, nodes := range result {
		fmt.Printf("[ ")
		for _, node := range nodes {
			fmt.Printf("%d ", node.val)
		}
		fmt.Printf("]\n")
	}
}
