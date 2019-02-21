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

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	sum := MaxSubArray(nums)
	fmt.Println(sum)
}

func TestCoinChange(t *testing.T) {
	coins := []int{1, 2, 5, 7, 10}
	res := CoinChange(coins, 14)
	fmt.Println(res)
}

func TestMinimumTotal(t *testing.T) {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	min := MinimumTotal(triangle)
	fmt.Println(min)
}
