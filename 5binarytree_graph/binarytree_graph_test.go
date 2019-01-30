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
	g := NewTreeNode(4)
	h := NewTreeNode(7)
	i := NewTreeNode(2)
	j := NewTreeNode(5)
	k := NewTreeNode(1)
	a.left = b
	a.right = c
	b.left = d
	d.left = h
	d.right = i
	c.left = e
	c.right = g
	g.left = j
	g.right = k

	result := PathSum(a, 22)
	fmt.Println(result)
}
