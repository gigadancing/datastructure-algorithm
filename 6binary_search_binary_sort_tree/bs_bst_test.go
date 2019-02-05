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
