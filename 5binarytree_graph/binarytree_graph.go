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

// 例2. 已知二叉树，求二叉树中给定的两个节点的最近公共祖先。
// 最近公共祖先：两节点v与w的最近公共祖先u，满足在树上最低（离根最远）且v、w都是u的子孙。
//                  3
//                 / \
//                5   1
//               / \ / \
//              6  2 0  8
//                / \
//               7   4
// 例如： 6、4的最近公共祖先是5。
// 思路：
// 1. 从根节点遍历至该节点后结束。
// 2. 将遍历过程中的节点按顺序存储起来。
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var ancestor *TreeNode
	path := make([]*TreeNode, 0)          // 存放搜索路径的临时栈
	pPath := make([]*TreeNode, 0)         // p的路径
	qPath := make([]*TreeNode, 0)         // q的路径
	finish := 0                           // 结束标志
	search(root, p, path, &pPath, finish) // 搜索p
	path = path[:0]                       // 清空path
	finish = 0                            // 重置结束标志
	search(root, q, path, &qPath, finish) // 搜索q
	pathLen := 0                          // 较短路径长度
	if len(pPath) < len(qPath) {
		pathLen = len(pPath)
	} else {
		pathLen = len(qPath)
	}
	for i := 0; i < pathLen; i++ { // 找最近相同祖先
		if pPath[i] == qPath[i] {
			ancestor = pPath[i]
		}
	}
	return ancestor
}

func search(current, target *TreeNode, path []*TreeNode, result *[]*TreeNode, finish int) {
	if current == nil || finish == 1 { // 当前节点为空或已找到节点，结束搜索
		return
	}
	path = append(path, current) // 当前节点加入到路径中
	if current == target {       // 搜索到目标节点，将目标节点加入路径，停止搜索（finish置为1）
		finish = 1
		*result = append(*result, path...)
	}
	search(current.left, target, path, result, finish)  // 搜索当前节点的左孩子
	search(current.right, target, path, result, finish) // 搜索当前节点的右孩子
	// 值传递不用删除，函数返回后自动就删除了
	//path = path[:len(path)-1]                           // 将当前节点从路径中删除
}
