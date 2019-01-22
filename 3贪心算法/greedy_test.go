package greedy

import (
	"fmt"
	"testing"
)

func TestFindContentChildren(t *testing.T) {
	g := MySlice{5, 10, 2, 9, 15, 9}
	s := MySlice{6, 1, 20, 3, 8}
	count := findContentChildren(s, g)
	fmt.Println("被满足的孩子个数:", count)
}

func TestWiggleMaxLength(t *testing.T) {
	nums := []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}
	length := WiggleMaxLength(nums)
	fmt.Println(length)
}

func TestRemoveKdigits(t *testing.T) {
	num := "1432219"
	str := RemoveKdigits(num, 3)
	fmt.Println(str)
}
