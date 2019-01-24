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

func TestRemoveKdigits2(t *testing.T) {
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
