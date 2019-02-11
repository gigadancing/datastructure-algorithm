package bg

import (
	"fmt"
	"testing"
)

func TestPathSum(t *testing.T) {
	a := NewTreeNode(5)
	b := NewTreeNode(4)
	c := NewTreeNode(8)
	d := NewTreeNode(11)
	e := NewTreeNode(13)
	f := NewTreeNode(4)
	g := NewTreeNode(7)
	h := NewTreeNode(2)
	i := NewTreeNode(5)
	j := NewTreeNode(1)
	a.Left = b
	a.Right = c
	b.Left = d
	d.Left = g
	d.Right = h
	c.Left = e
	c.Right = f
	f.Left = i
	f.Right = j

	result := PathSum(a, 22)
	fmt.Println(result)
}

func TestLowestCommonAncestor(t *testing.T) {
	a := NewTreeNode(3)
	b := NewTreeNode(5)
	c := NewTreeNode(1)
	d := NewTreeNode(6)
	e := NewTreeNode(2)
	f := NewTreeNode(0)
	g := NewTreeNode(8)
	h := NewTreeNode(7)
	i := NewTreeNode(4)
	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Left = f
	c.Right = g
	e.Left = h
	e.Right = i
	ancestor := LowestCommonAncestor(a, d, i)
	fmt.Printf("%d和%d的公共祖先是:%d\n", d.Val, i.Val, ancestor.Val)
	ancestor = LowestCommonAncestor(a, h, i)
	fmt.Printf("%d和%d的公共祖先是:%d\n", h.Val, i.Val, ancestor.Val)
	ancestor = LowestCommonAncestor(a, b, f)
	fmt.Printf("%d和%d的公共祖先是:%d\n", b.Val, f.Val, ancestor.Val)
	ancestor = LowestCommonAncestor(a, f, g)
	fmt.Printf("%d和%d的公共祖先是:%d\n", f.Val, g.Val, ancestor.Val)
	ancestor = LowestCommonAncestor(a, b, g)
	fmt.Printf("%d和%d的公共祖先是:%d\n", b.Val, g.Val, ancestor.Val)
	ancestor = LowestCommonAncestor(a, e, f)
	fmt.Printf("%d和%d的公共祖先是:%d\n", e.Val, f.Val, ancestor.Val)
}

func TestFlatten(t *testing.T) {
	a := NewTreeNode(3)
	b := NewTreeNode(5)
	c := NewTreeNode(1)
	d := NewTreeNode(6)
	e := NewTreeNode(2)
	f := NewTreeNode(0)
	g := NewTreeNode(8)
	h := NewTreeNode(7)
	i := NewTreeNode(4)
	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Left = f
	c.Right = g
	e.Left = h
	e.Right = i
	Flatten(a)

	for a != nil {
		fmt.Printf("%d ", a.Val)
		a = a.Right
	}
	fmt.Println()
}

func TestRightSideView(t *testing.T) {
	a := NewTreeNode(1)
	b := NewTreeNode(2)
	c := NewTreeNode(3)
	d := NewTreeNode(4)
	e := NewTreeNode(5)
	f := NewTreeNode(6)
	a.Left = b
	a.Right = c
	b.Right = e
	c.Right = d
	e.Left = f
	result := RightSideView(a)
	fmt.Printf("[ ")
	for _, node := range result {
		fmt.Printf("%d ", (*node).Val)
	}
	fmt.Printf("]\n")
}

func TestLeftSideView(t *testing.T) {
	a := NewTreeNode(1)
	b := NewTreeNode(2)
	c := NewTreeNode(3)
	d := NewTreeNode(4)
	e := NewTreeNode(5)
	f := NewTreeNode(6)
	a.Left = b
	a.Right = c
	b.Right = e
	c.Right = d
	e.Left = f
	result := LeftSideView(a)
	fmt.Printf("[ ")
	for _, node := range result {
		fmt.Printf("%d ", (*node).Val)
	}
	fmt.Printf("]\n")
}

func TestCanFinish(t *testing.T) {

}

func TestCanFinish2(t *testing.T) {

}
