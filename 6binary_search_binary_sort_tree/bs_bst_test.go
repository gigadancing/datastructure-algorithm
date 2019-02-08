package bst

import (
	"datastructure-algorithm/5binarytree_graph"
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

func TestBstPreorder(t *testing.T) {
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
	str := ""
	BstPreorder(a, &str)
	fmt.Println(str)
}

func TestSerializeAndDeserialize(t *testing.T) {
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
	res := Serialize(a)
	fmt.Println(res)
	root := Deserialize(res)
	bg.PreorderPrint(root, 0)
}
