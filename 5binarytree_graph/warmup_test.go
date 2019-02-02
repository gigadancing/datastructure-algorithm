package bg

import (
	"fmt"
	"github.com/eapache/queue"

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

func TestBFS(t *testing.T) {
	a := NewTreeNode(1)
	b := NewTreeNode(2)
	c := NewTreeNode(5)
	d := NewTreeNode(3)
	e := NewTreeNode(4)
	f := NewTreeNode(6)
	a.left = b
	a.right = c
	b.left = d
	b.right = e
	c.right = f

	result := Bfs(a)
	fmt.Printf("[ ")
	for _, n := range result {
		fmt.Printf("%d ", (*n).val)
	}
	fmt.Printf("]\n")
}

func TestBfsRecursion(t *testing.T) {
	a := NewTreeNode(1)
	b := NewTreeNode(2)
	c := NewTreeNode(5)
	d := NewTreeNode(3)
	e := NewTreeNode(4)
	f := NewTreeNode(6)
	a.left = b
	a.right = c
	b.left = d
	b.right = e
	c.right = f
	result := make([]*TreeNode, 0)
	q := queue.New()
	BfsRecursion(a, q, &result)
	fmt.Printf("[ ")
	for _, n := range result {
		fmt.Printf("%d ", (*n).val)
	}
	fmt.Printf("]\n")
}

func TestGraph(t *testing.T) {
	AdjacencyMatrix()
	AdjacencyList()
}
