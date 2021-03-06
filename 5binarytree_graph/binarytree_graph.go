package bg

import (
	"github.com/eapache/queue"
)

// 树节点
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 节点构造函数
func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
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
	pathValue += node.Val         // 更新路径和
	path = append(path, node.Val) // 将当前节点加入路径
	// 到达叶节点且路径和为sum，将路径加入到result
	if pathValue == sum && node.Left == nil && node.Right == nil {
		*result = append(*result, path)
	}
	preorder(node.Left, pathValue, sum, path, result)
	preorder(node.Right, pathValue, sum, path, result)
	pathValue -= node.Val
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
	search(current.Left, target, path, result, finish)  // 搜索当前节点的左孩子
	search(current.Right, target, path, result, finish) // 搜索当前节点的右孩子
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

	if node.Left == nil && node.Right == nil { // 叶节点
		*last = node
		return
	}

	// 备份左右子树
	left := node.Left
	right := node.Right

	// 左右子树的最后一个节点
	var (
		leftLast  *TreeNode
		rightLast *TreeNode
	)

	if left != nil { // 若有左子树，将左子树转换为单链表
		preorderTraverse(left, &leftLast)
		node.Left = nil
		node.Right = left
		*last = leftLast
	}

	if right != nil { // 若有右子树，将左子树转换为单链表
		preorderTraverse(right, &rightLast)
		if leftLast != nil {
			leftLast.Right = right
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

		if n.Left != nil {
			q.Add(NewPair(n.Left, depth+1))
		}

		if n.Right != nil {
			q.Add(NewPair(n.Right, depth+1))
		}
	}

	return views
}

// 从左侧观察
func LeftSideView(node *TreeNode) []*TreeNode {
	if node == nil {
		return nil
	}
	views := make([]*TreeNode, 0)
	q := queue.New()
	q.Add(NewPair(node, 0))

	for q.Length() != 0 {
		head := q.Peek().(*Pair)
		n := head.Node
		depth := head.Layer
		q.Remove()
		if depth == len(views) {
			views = append(views, n)
		}

		if n.Left != nil {
			q.Add(NewPair(n.Left, depth+1))
		}

		if n.Right != nil {
			q.Add(NewPair(n.Right, depth+1))
		}
	}

	return views
}

type CoursePair struct {
	Course int
	Depend int
}

func NewCoursePair(course, depend int) *CoursePair {
	return &CoursePair{
		Course: course,
		Depend: depend,
	}
}

// 例5. 课程安排
// 已知有n个课程，标记从0到n-1，课程之间是有依赖关系的。例如希望完成A课程，可能需要先完成B课程。已知n个课程的依赖关系，求是否可以将n个
// 课程全部完成。
// 课程依赖关系<课程1,课程2>代表课程1依赖课程2，如：
// [[1,0]]，返回true
// [[1,0],[0,1]]，返回false
// n个课程它们之间有m个依赖关系，可以看成顶点个数为n，边个数为m的有向图。故若为有向无环图，则可以完成全部课程，否则不能。
// 问题转换成构建图判断是否有环。
// 1. 深度优先搜索：如果正在搜索某一顶点（还未推出该顶点的递归深度搜索），又回到了该顶点证明有环。
// 2. 广度优先搜索（拓扑排序）：只将入度为0的节点添加至队列，当完成一个顶点的搜索（从队列中弹出），它指向的所有的顶点的入度都减1；若此时某个顶点的
// 入度为0则添加至队列。若完成广搜后所有顶点的入度都为0，则无环，否则有环。
func CanFinish(numCourses int, prerequisites []*CoursePair) bool {
	graph := make([]*GraphNode, 0)   // 邻接表
	visit := make([]int, numCourses) // 节点访问状态，-1代表没有访问过，0代表正在访问，1代表已经访问过
	// 创建图顶点并置访问状态为-1
	for i := 0; i < numCourses; i++ {
		graph = append(graph, NewGraphNode(i))
		visit[i] = -1
	}
	// 创建图
	for i := 0; i < len(prerequisites); i++ {
		// depend指向course
		begin := graph[prerequisites[i].Depend]
		end := graph[prerequisites[i].Course]
		begin.neighbors = append(begin.neighbors, end)
	}

	for i := 0; i < numCourses; i++ {
		// 若节点未访问过，则进行DFS；如果遇到环，则无法完成。
		if visit[i] == -1 && !DfsCourseGraph(graph[i], visit) {
			return false
		}

	}
	return true
}

// 返回值
// true - 无环
// false - 有环
func DfsCourseGraph(node *GraphNode, visit []int) bool {
	visit[node.label] = 0 // 状态置为0
	for i := 0; i < len(node.neighbors); i++ {
		if visit[node.neighbors[i].label] == -1 { // 未被访问
			if !DfsCourseGraph(node.neighbors[i], visit) {
				return false
			}
		} else if visit[node.neighbors[i].label] == 0 { // 正在访问，说明有环，返回false
			return false
		}
	}
	visit[node.label] = 1 // 状态置为1
	return true
}

// 广度优先搜索
func CanFinish2(numCourses int, prerequisites []*CoursePair) bool {
	graph := make([]*GraphNode, 0)
	degree := make([]int, 0)

	for i := 0; i < numCourses; i++ {
		degree = append(degree, 0)
		graph = append(graph, NewGraphNode(i))
	}

	for i := 0; i < len(prerequisites); i++ {
		begin := graph[prerequisites[i].Depend]
		end := graph[prerequisites[i].Course]
		begin.neighbors = append(begin.neighbors, end)
		degree[prerequisites[i].Course]++
	}

	q := queue.New()
	for i := 0; i < numCourses; i++ {
		if degree[i] == 0 {
			q.Add(graph[i])
		}
	}

	for q.Length() != 0 {
		front := q.Peek().(*GraphNode)
		q.Remove()
		for i := 0; i < len(front.neighbors); i++ {
			degree[front.neighbors[i].label]--
			if degree[front.neighbors[i].label] == 0 {
				q.Add(front.neighbors[i])
			}
		}
	}

	for _, d := range degree {
		if d != 0 { // 入度不为0
			return false
		}
	}

	return true
}
