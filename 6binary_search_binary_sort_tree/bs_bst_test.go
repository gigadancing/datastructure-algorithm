package bst

import (
	"fmt"
	"testing"
)

func TestBinarySearchRecursion(t *testing.T) {
	sortedArray := []int{-1, 2, 5, 20, 90, 100, 207, 800}
	find := BinarySearchRecursion(sortedArray, 0, len(sortedArray), 2)
	fmt.Println("2:", find)

	find = BinarySearchRecursion(sortedArray, 0, len(sortedArray), 200)
	fmt.Println("200:", find)
}
