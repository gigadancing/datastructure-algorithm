package bg

import (
	"testing"
)

func TestNewTreeNode(t *testing.T) {
	root := NewTreeNode(0)
	a := NewTreeNode(1)
	b := NewTreeNode(2)
	c := NewTreeNode(3)
	d := NewTreeNode(4)
	e := NewTreeNode(5)
	f := NewTreeNode(6)
	root.left = a
	root.right = b
	a.left = c
	a.right = d
	b.left = e
	b.right = f
	PreorderPrint(root, 0)

}
