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
