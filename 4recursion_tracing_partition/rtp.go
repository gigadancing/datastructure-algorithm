package rtp

import (
	"datastructure-algorithm/4recursion_tracing_partition/myslice"
	"github.com/golang-collections/collections/set"
	"sort"
)

// 1-a. 求子集
// 已知一组数（其中无重复元素），求这组数可以组成的所有子集。结果中不可有重复的子集。
// 例如：nums=[1,2,3]
// 结果为：[[],[1],[1,2],[1,2,3],[1,3],[2],[2,3]]
// 思路：
// 1. 递归
// nums = [1,2,3]，将子集[1]，[1,2]，[1,2,3]递归加入result
func SubsetsRecursion(i int, nums []int, item []int, result *[][]int) {
	if i >= len(nums) {
		return
	}
	item = append(item, nums[i])
	*result = append(*result, item)
	item = (item)[:len(item)-1]
	SubsetsRecursion(i+1, nums, item, result)
}

// 2. 位运算
// 若一个集合有ABC三个元素，则它的子集有2^3=8种组合，用0-7代表这些集合。
//    ABC  整数  集合
//    000   0    []
//    001   1    [C]
//    010   2    [B]
//    011   3    [B,C]
//    100   4    [A]
//    101   5    [A,C]
//    110   6    [A,B]
//    111   7    [A,B,C]
// 构造某一集合：使用A（100）、B（010）、C（001）与该集合对应的整数（000-111）做&运算，若结果为真，则放入item集合。
// 即：用某个集合代表的二进制与某个元素代表的的二进制做与运算，就可知道该元素是否在该集合中，也就能构造出该集合。
//    ABC  整数  集合           A是否出现          B是否出现            C是否出现
//    000   0    []          100 & 000 = 0     010 & 000 = 0       001 & 000 = 0
//    001   1    [C]         100 & 001 = 0     010 & 001 = 0       001 & 001 = 1
//    010   2    [B]         100 & 010 = 0     010 & 010 = 1       001 & 010 = 0
//    011   3    [B,C]       100 & 011 = 0     010 & 011 = 2       001 & 011 = 1
//    100   4    [A]         100 & 100 = 4     010 & 100 = 0       001 & 100 = 0
//    101   5    [A,C]       100 & 101 = 4     010 & 101 = 0       001 & 101 = 1
//    110   6    [A,B]       100 & 110 = 4     010 & 110 = 2       001 & 110 = 0
//    111   7    [A,B,C]     100 & 111 = 4     010 & 111 = 2       001 & 111 = 1
func SubsetsBitwiseOperation(nums []int) [][]int {
	result := make([][]int, 0)      // 结果集合
	allSets := 1 << uint(len(nums)) // 1<<n，2^n，设置全部集合的最大值+1
	for i := 0; i < allSets; i++ {
		item := make([]int, 0)
		// i代表从0至2^n-1这2^n个集合
		// i<<j代表nums数组的第j个元素，若i&(1<<j)为真，则放入入item
		for j := 0; j < len(nums); j++ {
			if i&(1<<uint(j)) != 0 { // 构造数字i代表的集合
				item = append(item, nums[j])
			}
		}
		result = append(result, item)
	}
	return result
}

// 1-b. 已知一数组（其中有重复元素），求这数组可以组成的所有子集（不可有重复的子集）。
// 例如：nums=[2,1,2,2]
// 结果为：[[],[1],[1,2],[1,2,2],[1,2,2,2],[2],[2,2],[2,2,2]]
// 注意：[2,1,2]与[1,2,2]是重复的集合
// 重复原因：
// 由于集合的元素是无序的。
// 1. 不同位置的元素组成的集合是同一个子集，顺序相同。
// 2. 不同位置的元素组成的集合是同一个子集，顺序不同。
func SubsetsWithDup(nums myslice.MySlice) [][]int {
	result := make([][]int, 0)
	item := make(myslice.MySlice, 0)
	resSet := set.New() // 用于去重的集合
	sort.Sort(&nums)    // 排序
	result = append(result, item)
	generateItem(0, nums, result, item, resSet)
	return result
}

func generateItem(i int, nums []int, result [][]int, item []int, resSet *set.Set) {
	if i >= len(nums) {
		return
	}
	item = append(item, nums[i])
	if !resSet.Has(item) { // set中无item
		result = append(result, item) // 将item加入结果数组
		resSet.Insert(item)           // 将item放入去重集合
	}
	generateItem(i+1, nums, result, item, resSet)
	item = item[:len(item)-1]
	generateItem(i+1, nums, result, item, resSet)
}

// 1-c. 组合数之和
// 已知一组数（其中有重复元素），求这组数可以组成的所有子集，子集中各元素之和为target的子集，结果中无重复的子集。
// 例如：nums=[10,1,2,7,6,1,5]，target=8
// 结果为：[[1,7],[1,2,5],[2,6],[1,1,6]]
func CombinationSum(nums myslice.MySlice, target int) [][]int {
	result := make([][]int, 0)
	item := make([]int, 0)
	resSet := set.New()
	sort.Sort(&nums)
	generateTargetItem(0, nums, result, item, resSet, 0, target)
	return result
}

func generateTargetItem(i int, nums []int, result [][]int, item []int, resSet *set.Set, sum int, target int) {
	// 当元素已经选完或sum的值超过target
	if i >= len(nums) || sum > target {
		return
	}
	sum += nums[i]
	item = append(item, nums[i])
	// item中元素和为target且item未在集合中
	if target == sum && !resSet.Has(item) {
		result = append(result, item)
		resSet.Insert(item)
	}
	generateTargetItem(i+1, nums, result, item, resSet, sum, target)
	// 回溯后，sum将nums[i]减去并从item中删除
	sum -= nums[i]
	item = item[:len(item)-1]
	generateTargetItem(i+1, nums, result, item, resSet, sum, target)
}

// 2. 生成括号
// 已知n组括号，开发一个程序，生成这n组括号所有合法的组合可能。
// 例如：n=3
// 结果为：["((()))", "(()())", "(())()", "()(())", "()()()"]
// 不考虑"("和")"的个数的情况下，生成所有的组合
// 字符串长度为2n，每个位置有2种选择，故共有2^2n=4^n种组合
func generateCombination(item string, n int, result *[]string) {
	// 当字符串长度是括号组数2倍时，递归结束
	if len(item) == 2*n {
		*result = append(*result, item)
		return
	}
	generateCombination(item+"(", n, result) // 添加"("，继续递归
	generateCombination(item+")", n, result) // 添加")"，继续递归
}

// 合法：
// 1. 在所有的可能中，左右括号的数量不会超过n（n组括号，即n个左括号，n个右括号）。
// 2. 放一个左括号才能放一个右括号，即右括号不可先于左括号放置。
// 递归限制条件：
// 1. 左右括号的数量最多n个。
// 2. 若左括号的数量<=右括号的数量，则不可进行放置右括号的递归。
// item - 当前生成的字符串
// left - 当前还可放左括号的数量
// right - 当前还可放右括号的数量
func generate(item string, left, right int, result *[]string) {
	if left == 0 && right == 0 {
		*result = append(*result, item)
		return
	}
	// 若可以放左括号，先放左括号
	if left > 0 {
		generate(item+"(", left-1, right, result)
	}
	// 剪枝：满足一定条件才放右括号
	// 剩余可放左括号数量<剩余可放右括号数量，即：已放左括号数量>已放右括号数量，说明可以放右括号
	if left < right {
		generate(item+")", left, right-1, result)
	}
}

//
func GenerateParenthesis(n int) []string {
	result := make([]string, 0)
	generate("", n, n, &result)
	return result
}
