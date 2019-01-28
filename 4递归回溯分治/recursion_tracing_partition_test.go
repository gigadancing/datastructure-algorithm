package rtp

import (
	"fmt"
	"testing"
)

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
