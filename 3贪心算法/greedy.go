package greedy

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
	"sort"
	"strconv"
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
// 已知一个使用字符串表示非负整数num，将num中的k个数组移除，求移除k个数字后，可以获得的最小的可能的新数字。
// （num不会以0开头，num的长度小于1002）
// 例如：
// 输入：num="1432219"，k=3
// 去掉三个数字后得到的可能很多，如1432，4322，2219，1219；去掉数字4，3，9后得到1219最小。
// 思路：
// 要让得到的数字最小，那么优先高位最小。
func RemoveKdigits(num string, k int) string {
	if k == 0 {
		return num
	}
	s := stack.New()
	result := ""
	for _, v := range num {
		n := int(v - '0')
		// 栈不为空，栈顶元素大于n，仍然可删
		for s.Len() > 0 && k > 0 && s.Peek().(int) > n {
			s.Pop()
			k--
		}
		// s.Len()>0是解决如X0XXX这种情况，说明0前面有非零数字，所以要将0加入
		// 即：当s不为空时，无论n是否为0都直接入栈；当s为空时，n>0才入栈。
		// 这里用了逻辑运算的短路特性：
		// s.Len()>0则n!=0不会运算；当s.Len()==0时，n!=0才运算。
		if s.Len() > 0 || n != 0 {
			s.Push(n) // 加入数据
		}
	}

	// 如果最后K>0，仍然可删
	if s.Len() > 0 && k > 0 {
		s.Pop()
		k--
	}
	// 将栈中的结果转成字符串存到result中
	for s.Len() > 0 {
		result = strconv.Itoa(s.Pop().(int)) + result
	}
	// 若result为空，则结果为0
	if result == "" {
		result = "0"
	}

	return result
}

// 4-a 跳跃游戏
// 一个数组存储了非负整型数据，数组中的第i个元素nums[i]，代表了可从数组第i个位置最多向前跳跃nums[i]步，已知数组各元素的情况下，求是否
// 可以从数组第0个位置跳跃至数组的最后一个位置。
// nums=[2,3,1,1,4]可以从nums[0]=2跳跃至nums[4]=4
// nums=[3,2,1,0,4]不可以从nums[0]=3跳跃至nums[4]=4
// 思路：
// 1. 求从第i个位置最远可以跳至第index[i]位置。根据第i位置最远可跳nums[i]步：index[i] = nums[i] + i。
// 2. 初始化jump为0，maxIndex为index[0]，maxIndex代表从0位置到jump位置这个过程中，最远可以达到的位置。
// 3. 利用jump扫描index数组，直到jump达到数组尾部或jump超过maxIndex为止，扫描过程中更新maxIndex。
// 4. 若最终jump为数组长度，则返回true，否则返回false。
func CanJump(nums []int) bool {
	length := len(nums)
	if length < 2 {
		return true
	}

	// 每个位置最远可跳的位置
	index := make([]int, length)
	for i, v := range nums {
		index[i] = v + i
	}

	// 初始化jump和maxIndex
	jump := 0
	maxIndex := index[0]

	// jump跳至数组尾部或jump超越了当前可跳跃的最远位置
	for jump < length && jump <= maxIndex {
		// 如果可以跳更远，则更新maxIndex
		if maxIndex < index[jump] {
			maxIndex = index[jump]
		}
		jump++
	}

	// jump调至数组尾部
	if jump == length {
		return true
	}

	return false
}

// 4-b 跳跃游戏2
// 一个数组存储了非负整型数据，数组中的第i个元素nums[i]，代表了可从数组第i个位置最多向前跳跃nums[i]步，已知数组各元素的情况下，确认
// 可以从第0个位置跳跃至数组最后一个位置，求最少需要跳跃几次。
// 例如：
// nums=[2,3,1,1,4]，从第0个位置跳到第一个位置，从第一个位置跳到最后一个位置。
// 思路：
// 1. 设置currentMaxIndex为当前可达到的最远位置。
// 2. 设置preMaxIndex为在遍历各个位置的过程中，各个位置可达到的最远位置。
// 3. 设置jumpMin为最少跳跃次数。
// 4. 利用i遍历nums数组，若i到达currentMaxIndex，则jumpMin加1，currentMaxIndex=preMaxIndex
// 5. 遍历过程中，若nums[i]+i(index[i])更大，则更新preMaxIndex=nums[i]+i
func Jump(nums []int) int {
	length := len(nums)
	if length < 2 { // 长度小于2不需跳跃
		return 0
	}

	// 当前可达到的最远位置
	currentMaxIndex := nums[0]
	// 遍历各个过程中可达到的最远位置
	preMaxIndex := nums[0]
	// 最小跳跃次数
	jumpMin := 0

	for i := 1; i < length; i++ {
		if i == currentMaxIndex {
			jumpMin++
			fmt.Println(i)
			currentMaxIndex = preMaxIndex
		}
		if preMaxIndex < nums[i]+i {
			preMaxIndex = nums[i] + 1
		}
	}
	return jumpMin
}

type Points [][]int

func (p *Points) Len() int {
	return len(*p)
}

func (p *Points) Less(i, j int) bool {
	return (*p)[i][0] < (*p)[j][0]
}

func (p *Points) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

// 5 射击气球
// 已知在一个平面上有一定数量的气球，平面可以看作一个坐标系，在平面的x轴的不同位置安排弓箭手向y轴方向射箭，弓箭可以向y轴走无穷远；
// 给定气球的宽度xstart <= x<= xend，问至少需要多少弓箭手将气球全部射爆？
// 例如：
// 四个气球[10,16] [2,8] [1,6] [7,12]，至少需要2个弓箭手
// 思路：尽可能让一支箭射穿更多的气球
// 1. 对各个气球进行排序，按照气球的左端点从小到大排序。
// 2. 遍历气球数组，同时维护一个设计区间，在满足可以将当前气球射穿的情况下，尽可能击穿更多的气球，更新一次射击区间（保证射击区间可以将新气球也击穿）。
// 3. 如果新气球没有办法击穿，则需增加一名射手，即维护一个新的设计区间将气球击穿，随后继续遍历气球数组。
func FindMinArrowShots(points Points) int {
	if len(points) == 0 {
		return 0
	}

	sort.Sort(&points) // 对气球按左端点从小到大排序
	shootNum := 1      // 初始化射手数量为1
	// 初始化射击区间，即第一个气球的两端点
	shootBegin := points[0][0]
	shootEnd := points[0][1]
	shootBegin = shootBegin
	for i := 1; i < len(points); i++ {
		if points[i][0] <= shootEnd {
			shootBegin = points[i][0]
			if shootEnd > points[i][1] {
				shootEnd = points[i][1]
			}
		} else {
			shootNum++
			shootBegin = points[i][0]
			shootEnd = points[i][1]
		}
	}

	return shootNum
}
