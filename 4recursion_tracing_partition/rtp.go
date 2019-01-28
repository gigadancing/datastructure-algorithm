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
	if resSet.Has(item) {
		result = append(result, item)
		resSet.Insert(item)
	}
	generateItem(i+1, nums, result, item, resSet)
	item = item[:len(item)-1]
	generateItem(i+1, nums, result, item, resSet)
}
