package bst

import "datastructure-algorithm/5binarytree_graph"

// 二分查找递归
// 返回值：
// true - 找到
// false - 未找到
func BinarySearchRecursion(sortedArray []int, begin, end, target int) bool {
	if len(sortedArray) == 0 || begin > end {
		return false
	}
	mid := (begin + end) / 2
	if target == sortedArray[mid] {
		return true
	} else if target < sortedArray[mid] {
		return BinarySearchRecursion(sortedArray, begin, mid-1, target)
	} else {
		return BinarySearchRecursion(sortedArray, mid+1, end, target)
	}
}

func BinarySearch(sortedArray []int, target int) bool {
	begin := 0
	end := len(sortedArray) - 1

	for begin <= end {
		mid := (begin + end) / 2
		if target == sortedArray[mid] {
			return true
		} else if target < sortedArray[mid] {
			end = mid - 1
		} else {
			begin = mid + 1
		}
	}
	return false
}

// 二叉排序树
// 它是一棵具有下列性质的二叉树：
// 1. 若左子树不空，则左子树上所有节点的值均小于或等于它的根节点的值
// 2. 若右子树不空，则右子树上所有节点的值均大于或等于它的根节点的值
// 3. 左右子树也分别为二叉排序树
// 4. 等于的情况只能出现在左子树或右子树的某一侧
// 二叉排序树插入节点
func BstInsert(node, insertNode *bg.TreeNode) {
	if insertNode.Val < node.Val { // 插入左子树
		if node.Left == nil { // 左子树为空，直接插入左子树
			node.Left = insertNode
		} else { // 递归插入左子树
			BstInsert(node.Left, insertNode)
		}
	} else { // 插入右子树
		if node.Right == nil { // 右子树为空，直接插入右子树
			node.Right = insertNode
		} else { // 递归插入右子树
			BstInsert(node.Right, insertNode)
		}
	}
}

// 二叉排序树查找值
func BstSearch(node *bg.TreeNode, value int) bool {
	if node.Val == value {
		return true
	}
	if value < node.Val { // 在左子树中查找
		if node.Left != nil {
			return BstSearch(node.Left, value)
		} else {
			return false
		}
	} else { // 在右子树中查找
		if node.Right != nil {
			return BstSearch(node.Right, value)
		} else {
			return false
		}
	}
}
