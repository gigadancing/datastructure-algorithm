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
	a.left = b
	a.right = c
	b.left = d
	d.left = g
	d.right = h
	c.left = e
	c.right = f
	f.left = i
	f.right = j

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
	a.left = b
	a.right = c
	b.left = d
	b.right = e
	c.left = f
	c.right = g
	e.left = h
	e.right = i
	ancestor := LowestCommonAncestor(a, d, i)
	fmt.Printf("%d和%d的公共祖先是:%d\n", d.val, i.val, ancestor.val)
	ancestor = LowestCommonAncestor(a, h, i)
	fmt.Printf("%d和%d的公共祖先是:%d\n", h.val, i.val, ancestor.val)
	ancestor = LowestCommonAncestor(a, b, f)
	fmt.Printf("%d和%d的公共祖先是:%d\n", b.val, f.val, ancestor.val)
	ancestor = LowestCommonAncestor(a, f, g)
	fmt.Printf("%d和%d的公共祖先是:%d\n", f.val, g.val, ancestor.val)
	ancestor = LowestCommonAncestor(a, b, g)
	fmt.Printf("%d和%d的公共祖先是:%d\n", b.val, g.val, ancestor.val)
	ancestor = LowestCommonAncestor(a, e, f)
	fmt.Printf("%d和%d的公共祖先是:%d\n", e.val, f.val, ancestor.val)
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
	a.left = b
	a.right = c
	b.left = d
	b.right = e
	c.left = f
	c.right = g
	e.left = h
	e.right = i
	Flatten(a)

	for a != nil {
		fmt.Printf("%d ", a.val)
		a = a.right
	}
	fmt.Println()
}

func TestRightsideView(t *testing.T) {
	a := NewTreeNode(1)
	b := NewTreeNode(2)
	c := NewTreeNode(3)
	d := NewTreeNode(4)
	e := NewTreeNode(5)
	f := NewTreeNode(6)
	a.left = b
	a.right = c
	b.right = e
	c.right = d
	e.left = f
	result := RightSideView(a)
	fmt.Printf("[ ")
	for _, node := range result {
		fmt.Printf("%d ", (*node).val)
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
	a.left = b
	a.right = c
	b.right = e
	c.right = d
	e.left = f
	result := LeftSideView(a)
	fmt.Printf("[ ")
	for _, node := range result {
		fmt.Printf("%d ", (*node).val)
	}
	fmt.Printf("]\n")
}
