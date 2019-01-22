package greedy

import (
	"fmt"
	"sort"
)

type MySlice []int

func (ms *MySlice) Len() int {
	return len(*ms)
}
func (ms *MySlice) Swap(i, j int) {
	(*ms)[i], (*ms)[j] = (*ms)[j], (*ms)[i]
}
func (ms *MySlice) Less(i, j int) bool {
	return (*ms)[i] < (*ms)[j]
}

// 1. 分糖果
// 已知一些孩子和一些糖果，每个孩子有需求因子g，每个糖果的大小s，当某个糖果的大小s>=某个孩子的需求因子g时，代表该糖果可以满足该孩子；
// 求使用这些糖果，最多能满足多少孩子？（注意，某个孩子最多只能用一个糖果满足）
// 思路：
// 对需求因子g和糖果大小s进行从小打到排序
// 按糖果大小顺序尝试是否满足某个孩子，每个糖果只尝试一次。若成功，则换下一个孩子；直到没有更过孩子或没有更多糖果为止。
func findContentChildren(s, g MySlice) int {
	sort.Sort(&g)
	sort.Sort(&s)
	index := 0 // 当前处理的孩子的位置
	children := make([]int, 0)
	for _, v := range s {
		if v >= g[index] && index < len(g) {
			children = append(children, g[index])
			index++
		}
	}
	for _, v := range children {
		fmt.Printf("被满足的孩子 %d\n", v)
	}
	return len(children)
}

// 2. 摇摆序列
// 一个整数序列，如果两个相邻元素的差恰好正负（负正）交替出现，则该序列被称为摇摆序列。
// 一个小于等于2个元素的序列直接为摇摆序列。
// 例如：
// 序列[1,7,4,9,2,5]，相邻元素的差(6,-3,5,-7,3)，该序列为摇摆序列。
// 序列[1,4,7,2,5](3, 3, -3, 3)、[1,7,4,5,5](6,-3,1,0)不是摇摆序列。
// 给定一个随机序列，求这个序列满足的摇摆序列定义的最长子序列的长度。
// 例如：
// 输入[1,7,4,9,2,5]，结果为6；输入[1,17,5,10,13,15,10,5,16,8]结果为7([1,17,10,13,10,16,8])；输入[1,2,3,4,5,6,7,8,9]结果为2。
// 		                        BEGIN
// 		                      /       \
// 		                     /         \
//	       nums[i-1]<nums[i]/           \ nums[i-1]>nums[i]
// 		                   /             \
// 		                  /               \
// 		                 /                 \
// 		               |/ nums[i-1]>nums[i] \|
// 	(nums[i-1]<nums[i])UP------------------->DOWN(nums[i-1]>nums[i])
// 	              	     <-------------------
//  	                   nums[i-1] < nums[i]
// 状态机三种状态
const (
	BEGIN = iota // 开始
	UP           // 上升
	DOWN         // 下降
)

func WiggleMaxLength(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	state := BEGIN // 初始状态
	maxLenght := 2 // 摇摆序列长度至少为2

	for i := 2; i < len(nums); i++ {
		switch state {
		case BEGIN:
			if nums[i-1] < nums[i] {
				state = UP
				maxLenght++
			} else if nums[i-1] > nums[i] {
				state = DOWN
				maxLenght++
			}
		case UP:
			if nums[i-1] > nums[i] {
				state = DOWN
				maxLenght++
			}
		case DOWN:
			if nums[i-1] < nums[i] {
				state = UP
				maxLenght++
			}
		}
	}
	return maxLenght
}

// 3. 移除K个数字
// 已知一个使用字符串表示非负整数num，将num中的k个数组移除，求移除k个数字后，可以获得的最小的可能的新数字。（num不会以0开头，num的长度
// 小于1002）
// 例如：
// 输入：num="1432219"，k=3
// 去掉三个数字后得到的可能很多，如1432，4322，2219，1229，1221；去掉数字4，3，9后得到1221最小。
func RemoveKdigits(num string, k int) string {

	return ""
}
