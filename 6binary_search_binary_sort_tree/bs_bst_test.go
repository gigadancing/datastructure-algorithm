package bst

import (
	"fmt"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	sortedArray := []int{1, 3, 5, 6}
	target := []int{0, 1, 2, 3, 4, 5, 6, 7}
	for _, t := range target {
		index := SearchInsert(sortedArray, t)
		fmt.Println(t, "-->", index)
	}
}

func TestSearchRange(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 8, 8, 10}
	for i := 0; i < 12; i++ {
		res := SearchRange(nums, i)
		fmt.Printf("%d:[%d,%d]\n", i, res[0], res[1])
	}
}

func TestSearch(t *testing.T) {
	nums := []int{9, 12, 15, 20, 1, 3, 6, 7}
	for i := 0; i < 22; i++ {
		fmt.Printf("%d:%d\n", i, Search(nums, i))
	}
}
