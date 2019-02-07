package bst

import (
	"datastructure-algorithm/5binarytree_graph"
	"fmt"
	"testing"
)

func TestBinarySearchRecursion(t *testing.T) {
	sortedArray := []int{-1, 2, 5, 20, 90, 100, 207, 800}
	find := BinarySearchRecursion(sortedArray, 0, len(sortedArray)-1, 2)
	fmt.Println("2:", find)

	find = BinarySearchRecursion(sortedArray, 0, len(sortedArray)-1, 200)
	fmt.Println("200:", find)
}

func TestBinarySearch(t *testing.T) {
	sortedArray := []int{-1, 2, 5, 20, 90, 100, 207, 800}
	find := BinarySearch(sortedArray, 200)
	fmt.Println("200:", find)

	find = BinarySearch(sortedArray, 2)
	fmt.Println("2:", find)
}

func TestBstInsert(t *testing.T) {
	root := bg.NewTreeNode(8)
	test := []int{3, 10, 1, 6, 15}
	for _, v := range test {
		BstInsert(root, bg.NewTreeNode(v))
	}
	bg.PreorderPrint(root, 0)
}

func TestBstSearch(t *testing.T) {
	a := bg.NewTreeNode(8)
	b := bg.NewTreeNode(3)
	c := bg.NewTreeNode(10)
	d := bg.NewTreeNode(1)
	e := bg.NewTreeNode(6)
	f := bg.NewTreeNode(15)
	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f
	for i := 0; i < 20; i++ {
		if BstSearch(a, i) {
			fmt.Printf("%d is in the BST.\n", i)
		} else {
			fmt.Printf("%d is not in the BST.\n", i)
		}
	}
}
