package bg

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		val:   val,
		left:  nil,
		right: nil,
	}
}

// 例1：路径之和
// 给定一个二叉树与整数sum，找出所有从根节点到叶节点的路径，这些路径上的节点值累加和为sum。
//           5
//          / \
//         4   8
//        /   / \
//       11  13  4
//      / \     / \
//     7   2   5   1
// 例如： sum=22
// 结果为：[[5,4,11,2],[5,8,4,5]]
// 思路：
// 深度搜索所有从根节点到叶节点的路径，检查个路径上所有节点的值的和是否为sum。
// 1.从根节点深度遍历二叉树，先序遍历时，将该节点值存储至path栈中，使用pathValue累加节点值。
// 2.当遍历至叶节点时，检查pathValue值是否为sum，若为sum，则将path加入result结果中。
// 3.在后序遍历时，将该节点从path栈中弹出，pathValue减去节点值。
func PathSum(root *TreeNode, sum int) (result [][]int) {
	if root == nil {
		return
	}
	path := make([]int, 0)
	pathValue := 0
	preorder(root, pathValue, sum, path, &result)
	return
}

// 先序遍历
func preorder(node *TreeNode, pathValue, sum int, path []int, result *[][]int) {
	if node == nil {
		return
	}
	pathValue += node.val         // 更新路径和
	path = append(path, node.val) // 将当前节点加入路径
	// 到达叶节点且路径和为sum，将路径加入到result
	if pathValue == sum && node.left == nil && node.right == nil {
		*result = append(*result, path)
	}
	preorder(node.left, pathValue, sum, path, result)
	preorder(node.right, pathValue, sum, path, result)
	pathValue -= node.val
	path = path[:len(path)-1] // 将当前节点从路径中删除
	return
}
