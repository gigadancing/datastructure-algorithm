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
