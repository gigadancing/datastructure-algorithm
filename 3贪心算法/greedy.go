package greedy

import (
	"fmt"
	"sort"
)

// 1. 分糖果
// 已知一些孩子和一些糖果，每个孩子有需求因子g，每个糖果的大小s，当某个糖果的大小s>=某个孩子的需求因子g时，代表该糖果可以满足该孩子；
// 求使用这些糖果，最多能满足多少孩子？（注意，某个孩子最多只能用一个糖果满足）

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
