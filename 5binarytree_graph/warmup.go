package bg

import (
	"fmt"
	"github.com/eapache/queue"
)

// 先序打印二叉树
func PreorderPrint(node *TreeNode, layer int) {
	if node == nil {
		return
	}
	for i := 0; i < layer; i++ {
		fmt.Printf("----")
	}
	fmt.Println(node.val)
	PreorderPrint(node.left, layer+1)
	PreorderPrint(node.right, layer+1)
}

// 遍历二叉树，记录从根节点到叶节点的所有路径
func Traverse(node *TreeNode, path []*TreeNode, result *[][]*TreeNode) {
	if node == nil {
		return
	}
	path = append(path, node)                  // 将当前节点加入路径
	if node.left == nil && node.right == nil { // 叶节点
		*result = append(*result, path)
	}
	Traverse(node.left, path, result)
	Traverse(node.right, path, result)
}

// 广度优先搜索
func Bfs(node *TreeNode) []*TreeNode {
	if node == nil {
		return nil
	}
	result := make([]*TreeNode, 0)
	q := queue.New()
	q.Add(node)
	for q.Length() != 0 {
		front := q.Peek().(*TreeNode)
		result = append(result, front)
		q.Remove()
		if front.left != nil {
			q.Add(front.left)
		}
		if front.right != nil {
			q.Add(front.right)
		}
	}
	return result
}

// 广度优先搜索递归实现
func BfsRecursion(node *TreeNode, q *queue.Queue, result *[]*TreeNode) {
	if node == nil {
		return
	}
	q.Add(node)
	if q.Length() == 0 {
		return
	}
	front := q.Remove().(*TreeNode)
	*result = append(*result, front)
	BfsRecursion(front.left, q, result)
	BfsRecursion(front.right, q, result)
}

const (
	MAXN = 5
)

// 图的构造与表示
// 1. 邻接矩阵
// 行代表顶点
// 列代表该顶点与其他顶点的关系
func AdjacencyMatrix() {
	graph := [MAXN][MAXN]int{} // 邻接矩阵
	graph[0][2] = 1
	graph[0][4] = 1
	graph[1][0] = 1
	graph[1][2] = 1
	graph[2][3] = 1
	graph[3][4] = 1
	graph[4][3] = 1
	for i := 0; i < MAXN; i++ {
		for j := 0; j < MAXN; j++ {
			fmt.Printf("%d ", graph[i][j])
		}
		fmt.Printf("\n")
	}
}

// 邻接表图节点
type GraphNode struct {
	label     int          // 顶点值
	neighbors []*GraphNode // 相邻节点指针数组
}

func NewGraphNode(x int) *GraphNode {
	return &GraphNode{
		label: x,
	}
}

// 2. 邻接表
func AdjacencyList() {
	graph := [MAXN]*GraphNode{}
	for i := 0; i < MAXN; i++ {
		graph[i] = NewGraphNode(i)
	}
	graph[0].neighbors = append(graph[0].neighbors, graph[2])
	graph[0].neighbors = append(graph[0].neighbors, graph[4])
	graph[1].neighbors = append(graph[1].neighbors, graph[0])
	graph[1].neighbors = append(graph[1].neighbors, graph[2])
	graph[2].neighbors = append(graph[2].neighbors, graph[3])
	graph[3].neighbors = append(graph[3].neighbors, graph[4])
	graph[4].neighbors = append(graph[4].neighbors, graph[3])

	for i := 0; i < MAXN; i++ {
		fmt.Printf("Label(%d):", i)
		for j := 0; j < len(graph[i].neighbors); j++ {
			fmt.Printf(" %d ", graph[i].neighbors[j].label)
		}
		fmt.Printf("\n")
	}

}
