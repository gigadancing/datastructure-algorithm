package biq

import (
	"fmt"
	"testing"
)

func TestFindAllEquivalentPoints(t *testing.T) {
	nums := []Point{{1, 1}, {1, 4}, {3, 4}, {3, 7}, {9, 7}, {9, 1}}
	res := FindAllEquivalentPoints(nums, 4)
	for _, p := range res {
		fmt.Println("(", p.x, ",", p.y, ")")
	}
}
