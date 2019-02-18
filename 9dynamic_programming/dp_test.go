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
