package rtp

import (
	"fmt"
	"testing"
)

func TestSubsetsRecursion(t *testing.T) {
	nums := []int{1, 2, 3}
	result := make([][]int, 0)
	item := make([]int, 0)
	SubsetsRecursion(0, nums, item, &result)
	fmt.Println(result)
}

func TestSubsetsBitwiseOperation(t *testing.T) {
	nums := []int{1, 2, 3}
	subsets := SubsetsBitwiseOperation(nums)
	fmt.Println(subsets)
}

func TestSubsetsWithDup(t *testing.T) {
	nums := []int{2, 1, 2, 2}
	subsets := SubsetsWithDup(nums)
	fmt.Println(subsets)
}
