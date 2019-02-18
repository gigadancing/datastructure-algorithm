package dp

import (
	"fmt"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	res := ClimbStairs(45)
	fmt.Println(res)
}

func TestClimbStairsDP(t *testing.T) {
	res := ClimbStairsDP(45)
	fmt.Println(res)
}

func TestRob(t *testing.T) {
	nums := []int{5, 2, 6, 3, 1, 7}
	res := Rob(nums)
	fmt.Println(res)
}
