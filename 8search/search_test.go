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
