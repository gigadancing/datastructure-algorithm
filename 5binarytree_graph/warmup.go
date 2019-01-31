package bg

import "fmt"

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
