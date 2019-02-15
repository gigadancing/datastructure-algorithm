package _search

import (
	"fmt"
	"testing"
)

func TestNumIslandsDFS(t *testing.T) {
	grid := [][]int{
		{1, 1, 1, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0}}
	mark := NumIslandsDFS(grid)
	fmt.Println(mark)
}
