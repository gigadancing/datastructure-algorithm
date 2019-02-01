package bg

import "github.com/eapache/queue"

// 树节点
type TreeNode struct {
	val         int
	left, right *TreeNode
}

// 节点构造函数
func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		val:   val,
		left:  nil,
		right: nil,
	}
}

// 将节点和层数绑定的Pair
type Pair struct {
	Node  *TreeNode // 树节点
	Layer int       // 层数
}

// Pair构造函数
func NewPair(node *TreeNode, layer int) *Pair {
	return &Pair{
		Node:  node,
		Layer: layer,
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

// 例3. 二叉树转链表
// 给定一棵二叉树，将该二叉树就地（in-place）转为单链表。单链表中的节点顺序为二叉树的前序遍历顺序。
// 用节点的right来连接下一个节点，left为空。
func Flatten(root *TreeNode) {
	if root == nil {
		return
	}
	var last *TreeNode
	preorderTraverse(root, &last)
	return
}

func preorderTraverse(node *TreeNode, last **TreeNode) {
	if node == nil {
		return
	}

	if node.left == nil && node.right == nil { // 叶节点
		*last = node
		return
	}

	// 备份左右子树
	left := node.left
	right := node.right

	// 左右子树的最后一个节点
	var (
		leftLast  *TreeNode
		rightLast *TreeNode
	)

	if left != nil { // 若有左子树，将左子树转换为单链表
		preorderTraverse(left, &leftLast)
		node.left = nil
		node.right = left
		*last = leftLast
	}

	if right != nil { // 若有右子树，将左子树转换为单链表
		preorderTraverse(right, &rightLast)
		if leftLast != nil {
			leftLast.right = right
		}
		*last = rightLast
	}
}

// 例4. 给定一棵二叉树，假设从该二叉树的右侧观察它，将观察到的节点从上到下输出。
//          1
//         / \
//        2   3
//         \   \
//          5   4
//         /
//        6
// 结果为：[1,3,4,6]
// 思路：
// 对树进行广度优先搜索，将每一层最后一个节点加入结果集合。
// 将节点与层数绑定为pair，压入队列时，将节点和层数同时压入队列，并记录每一层出现的最后一个节点。
func RightSideView(node *TreeNode) []*TreeNode {
	if node == nil {
		return nil
	}

	views := make([]*TreeNode, 0) // 观察结果
	q := queue.New()              // 临时队列

	q.Add(NewPair(node, 0))

	for q.Length() != 0 {
		front := q.Peek().(*Pair) // 队列最前面的Pair
		depth := front.Layer      // 搜索的层数
		n := front.Node           // 搜索的节点
		q.Remove()                // 队列弹出一个Pair

		if len(views) == depth { // views的元素个数和层数相等，更新对应的层数为下标的元素；否则，加入元素
			views = append(views, n)
		} else {
			views[depth] = n
		}

		if n.left != nil {
			q.Add(NewPair(n.left, depth+1))
		}

		if n.right != nil {
			q.Add(NewPair(n.right, depth+1))
		}
	}

	return views
}
