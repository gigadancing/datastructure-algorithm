package _search

import (
	"container/heap"
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
	mark, num := NumIslandsDFS(grid)
	fmt.Println(mark)
	fmt.Println(num)
}

func TestNumIslandsBFS(t *testing.T) {
	grid := [][]int{
		{1, 1, 1, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0}}
	mark := NumIslandsBFS(grid)
	fmt.Println(mark)
}

func TestLadderLength(t *testing.T) {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	steps := LadderLength(beginWord, endWord, wordList)
	fmt.Println(steps)
	wordList2 := []string{"hot", "lot", "low", "cow", "cog", "dog", "dot"}
	steps = LadderLength(beginWord, endWord, wordList2)
	fmt.Println(steps)
}

func TestFindLadders(t *testing.T) {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	ladders := FindLadders(beginWord, endWord, wordList)
	for _, ladder := range ladders {
		fmt.Println(ladder)
	}
}

func TestMakeSquare(t *testing.T) {
	nums := []int{1, 1, 2, 2, 2}
	ok := MakeSquare(nums)
	fmt.Println(ok)
	nums = []int{3, 3, 4, 4, 4}
	ok = MakeSquare(nums)
	fmt.Println(ok)
	nums = []int{1, 1, 2, 4, 3, 2, 3}
	ok = MakeSquare(nums)
	fmt.Println(ok)
	nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 5, 4, 3, 2, 1}
	ok = MakeSquare(nums)
	fmt.Println(ok)
}

func TestQueueItemSlice(t *testing.T) {
	q := make(QueueItemSlice, 0)
	heap.Init(&q)
	heap.Push(&q, NewQueueItem(0, 0, 5))
	heap.Push(&q, NewQueueItem(1, 3, 2))
	heap.Push(&q, NewQueueItem(5, 2, 4))
	heap.Push(&q, NewQueueItem(0, 1, 8))
	heap.Push(&q, NewQueueItem(6, 7, 1))

	for q.Len() != 0 {
		item := heap.Pop(&q).(*QueueItem)
		fmt.Printf("(%d,%d,%d)\n", item.X, item.Y, item.Height)
	}
}

func TestTrapRainWater(t *testing.T) {
	heightMap := [][]int{{1, 4, 3, 1, 3, 2}, {3, 2, 1, 3, 2, 4}, {2, 3, 3, 2, 3, 1}}
	v := TrapRainWater(heightMap)
	fmt.Println(v)
}
