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
	fmt.Println("---------", num, "---------")
	res := RemoveKdigits(num, 3)
	fmt.Printf("去掉3个数字的最小值:%s\n", res)
	res = RemoveKdigits(num, 5)
	fmt.Printf("去掉5个数字的最小值:%s\n", res)
	res = RemoveKdigits(num, 1)
	fmt.Printf("去掉1个数字的最小值:%s\n", res)

	num = "100200"
	fmt.Println("---------", num, "---------")
	res = RemoveKdigits(num, 1)
	fmt.Printf("去掉1个数字的最小值:%s\n", res)
	res = RemoveKdigits(num, 2)
	fmt.Printf("去掉2个数字的最小值:%s\n", res)
	res = RemoveKdigits(num, 3)
	fmt.Printf("去掉3个数字的最小值:%s\n", res)
}

func TestCanJump(t *testing.T) {
	nums1 := []int{2, 3, 1, 1, 4} // true
	nums2 := []int{3, 2, 1, 0, 4} // false
	c1 := CanJump(nums1)
	c2 := CanJump(nums2)
	fmt.Println("nums1:", c1)
	fmt.Println("nums2:", c2)
}

func TestJump(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	count := Jump(nums)
	fmt.Println("最小跳跃次数:", count)
}
