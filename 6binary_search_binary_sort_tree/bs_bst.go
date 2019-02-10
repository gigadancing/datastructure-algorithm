package bst

import (
	"datastructure-algorithm/5binarytree_graph"
	"strconv"
	"strings"
)

// 例1. 插入位置
// 给定一个排序数组nums（无重复元素）与目标值target，如果target在nums里出现，则返回target所在下标，如果target在nums里未出现，则返回
// target应该插入位置的数组下标，使得将target插入数组nums后，数组仍有序。
func SearchInsert(nums []int, target int) int {
	index := -1
	begin := 0
	end := len(nums) - 1

	for index == -1 {
		mid := (begin + end) / 2
		if target == nums[mid] {
			index = mid
		} else if target < nums[mid] {
			if mid == 0 || target > nums[mid-1] {
				index = mid
			}
			end = mid - 1
		} else {
			if mid == len(nums)-1 || target < nums[mid-1] {
				index = mid + 1
			}
			begin = mid + 1
		}
	}
	return index
}

// 例2. 区间查找
// 给定一个排序数组nums（nums中有）与目标值target，如果target在nums中出现，则返回target所在区间的左右端点下标[左端点,右端点]，如果
// target在nums中未出现，则返回[-1,-1]
// 例如：
// nums=[5,7,7,8,8,8,8,10]
// target=8，返回：[3,6]
// target=6，返回：[-1,-1]
func SearchRange(nums []int, target int) []int {
	res := make([]int, 0)
	left := LeftBound(nums, target)
	right := RightBound(nums, target)
	res = append(res, left)
	res = append(res, right)
	return res
}

func LeftBound(nums []int, target int) int {
	begin := 0
	end := len(nums) - 1
	for begin <= end {
		mid := (begin + end) / 2
		if target == nums[mid] {
			if mid == 0 || nums[mid-1] < target {
				return mid
			}
			end = mid - 1
		} else if target < nums[mid] {
			end = mid - 1
		} else {
			begin = mid + 1
		}
	}
	return -1
}

func RightBound(nums []int, target int) int {
	begin := 0
	end := len(nums) - 1

	for begin <= end {
		mid := (begin + end) / 2
		if target == nums[mid] {
			if mid == len(nums)-1 || target < nums[mid+1] {
				return mid
			}
			begin = mid + 1
		} else if target < nums[mid] {
			end = mid - 1
		} else {
			begin = mid + 1
		}
	}
	return -1
}

// 例3. 旋转数组
// 给定一个排序数组nums（nums中有无重复元素），且nums可能以某个未知的下标旋转，给定目标值target，求target是否在nums中出现，若出现返回
// 所在下标，未出现返回-1。
// 例如：
// 原数组：nums=[1,3,6,7,9,12,15,20]
// 可能的旋转结果：
// [3,6,7,9,12,15,20,1]  [12,15,20,1,3,6,7,9]
// [6,7,9,12,15,20,1,3]  [15,20,1,3,6,7,9,12]
// [7,9,12,15,20,1,3,6]  [20,1,3,6,7,9,12,15]
// [9,12,15,20,1,3,6,7]
func Search(nums []int, target int) int {
	begin := 0
	end := len(nums) - 1
	for begin <= end {
		mid := (begin + end) / 2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			if nums[begin] < nums[mid] {
				if target >= nums[begin] {
					end = mid - 1
				} else {
					begin = mid + 1
				}
			} else if nums[begin] > nums[mid] {
				end = mid - 1
			} else {
				begin = mid + 1
			}
		} else {
			if nums[begin] < nums[mid] {
				begin = mid + 1
			} else if nums[begin] > nums[mid] {
				if target >= nums[begin] {
					end = mid - 1
				} else {
					begin = mid + 1
				}
			} else {
				begin = mid + 1
			}
		}
	}
	return -1
}

// 例4.编码与解码
// 给定一棵二叉排序树，实现对该二叉排序树编码与解码功能。编码即将二叉排序树转为字符串，解码即将字符串转为二叉排序树。不限制使用何种编码
// 方法，只需保证当对二叉排序树调用编码功能后再调用解码功能将其复原。
func Serialize(node *bg.TreeNode) string {
	if node == nil {
		return ""
	}
	data := ""
	BstPreorder(node, &data)
	return data
}

func Deserialize(data string) *bg.TreeNode {
	res := strings.Split(data, "#")
	if len(res) == 0 {
		return nil
	}
	val, _ := strconv.Atoi(res[0])
	root := bg.NewTreeNode(val)
	for i := 1; i < len(res); i++ {
		if res[i] != "" {
			integer, _ := strconv.Atoi(res[i])
			BstInsert(root, bg.NewTreeNode(integer))
		}
	}
	return root
}

func BstPreorder(node *bg.TreeNode, data *string) {
	if node == nil {
		return
	}
	strVal := strconv.Itoa(node.Val) // 将节点值转为字符串
	*data += strVal + "#"
	BstPreorder(node.Left, data)
	BstPreorder(node.Right, data)
}

type BSTNode struct {
	Val         int
	Count       int // 左子树节点个数
	Left, Right *BSTNode
}

func NewBSTNode(value int) *BSTNode {
	return &BSTNode{
		Val: value,
	}
}

// 例5. 逆序数
// 已知数组nums，求新数组count，count[i]代表了在nums[i]右侧且比nums[i]小的元素个数。
// 例如：
// nums=[5,2,6,1],count=[2,1,1,0]
// nums=[6,6,6,1,1,1],count=[3,3,3,0,0,0]
// nums=[5,-7,9,1,3,5,-2,1],count=[5,0,5,1,2,2,0,0]
func CountSmaller(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	counts := make([]int, 0)     // 从后向前插入过程中比当前小的节点个数
	nodes := make([]*BSTNode, 0) // 二叉排序树节点
	// 逆序创建二叉排序树节点
	for i := len(nums) - 1; i >= 0; i-- {
		nodes = append(nodes, NewBSTNode(nums[i]))
	}
	counts = append(counts, 0) // 第一个的个数为0
	for i := 1; i < len(nodes); i++ {
		smallCount := 0
		InsertForInversionNumber(nodes[0], nodes[i], &smallCount)
		counts = append(counts, smallCount)
	}
	// 将counts逆序，还原为原来的顺序
	for i, j := 0, len(counts)-1; i < j; i, j = i+1, j-1 {
		counts[i], counts[j] = counts[j], counts[i]
	}
	return counts
}

func InsertForInversionNumber(node, insertedNode *BSTNode, smallerCount *int) {
	if insertedNode.Val <= node.Val {
		node.Count++
		if node.Left != nil {
			InsertForInversionNumber(node.Left, insertedNode, smallerCount)
		} else {
			node.Left = insertedNode
		}
	} else {
		*smallerCount++
		if node.Right != nil {
			InsertForInversionNumber(node.Right, insertedNode, smallerCount)
		} else {
			node.Right = insertedNode
		}
	}
}
