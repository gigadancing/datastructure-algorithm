package rtp

import (
	"fmt"
	"testing"
)

func TestSubSets(t *testing.T) {
	nums := []int{1, 2, 3}
	subsets := SubSetsBitwiseOperation(nums)
	fmt.Println(subsets)
}
