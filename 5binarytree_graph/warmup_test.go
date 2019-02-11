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
	root.Left = a
	root.Right = b
	a.Left = c
	a.Right = d
	b.Left = e
	b.Right = f
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
	root.Left = a
	root.Right = b
	a.Left = c
	a.Right = d
	b.Left = e
	b.Right = f
	d.Left = g
	e.Right = h
	f.Left = i
	f.Right = j
	path := make([]*TreeNode, 0)
	result := make([][]*TreeNode, 0)
	Traverse(root, path, &result)
	for _, nodes := range result {
		fmt.Printf("[ ")
		for _, node := range nodes {
			fmt.Printf("%d ", node.Val)
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
	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f

	result := Bfs(a)
	fmt.Printf("[ ")
	for _, n := range result {
		fmt.Printf("%d ", (*n).Val)
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
	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f
	result := make([]*TreeNode, 0)
	q := queue.New()
	BfsRecursion(a, q, &result)
	fmt.Printf("[ ")
	for _, n := range result {
		fmt.Printf("%d ", (*n).Val)
	}
	fmt.Printf("]\n")
}

func TestGraph(t *testing.T) {
	AdjacencyMatrix()
	AdjacencyList()
}

func TestDfsGraph(t *testing.T) {
	graph := [5]*GraphNode{}
	for i := 0; i < 5; i++ {
		graph[i] = NewGraphNode(i)
	}
	graph[0].neighbors = append(graph[0].neighbors, graph[4])
	graph[0].neighbors = append(graph[0].neighbors, graph[2])
	graph[1].neighbors = append(graph[1].neighbors, graph[0])
	graph[1].neighbors = append(graph[1].neighbors, graph[2])
	graph[2].neighbors = append(graph[2].neighbors, graph[3])
	graph[3].neighbors = append(graph[3].neighbors, graph[4])
	graph[4].neighbors = append(graph[4].neighbors, graph[3])
	visit := make([]int, MAXN)
	for i := 0; i < MAXN; i++ {
		if visit[i] == 0 {
			fmt.Printf("From label(%d) : ", graph[i].label)
			DfsGraph(graph[i], visit)
			fmt.Printf("\n")
		}
	}
}

func TestBfsGraph(t *testing.T) {
	graph := [5]*GraphNode{}
	for i := 0; i < 5; i++ {
		graph[i] = NewGraphNode(i)
	}
	graph[0].neighbors = append(graph[0].neighbors, graph[4])
	graph[0].neighbors = append(graph[0].neighbors, graph[2])
	graph[1].neighbors = append(graph[1].neighbors, graph[0])
	graph[1].neighbors = append(graph[1].neighbors, graph[2])
	graph[2].neighbors = append(graph[2].neighbors, graph[3])
	graph[3].neighbors = append(graph[3].neighbors, graph[4])
	graph[4].neighbors = append(graph[4].neighbors, graph[3])
	visit := make([]int, MAXN)
	for i := 0; i < MAXN; i++ {
		if visit[i] == 0 {
			fmt.Printf("From label(%d) : ", graph[i].label)
			BfsGraph(graph[i], visit)
			fmt.Printf("\n")
		}
	}
}
