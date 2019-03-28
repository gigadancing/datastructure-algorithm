package string

import (
	"fmt"
	"testing"
)

func TestShiftLetters(t *testing.T) {
	S := "abc"
	shifts := []int{3, 5, 9}
	res := shiftingLetters(S, shifts)
	fmt.Println(res)
}

func TestWordBreak(t *testing.T) {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	fmt.Println(wordBreak(s, wordDict))
	s = "applepenapple"
	wordDict = []string{"apple", "pen"}
	fmt.Println(wordBreak(s, wordDict))
	s = "catsandog"
	wordDict = []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println(wordBreak(s, wordDict))
}

func TestDifferentWaysToAddParentheses(t *testing.T) {
	input := "2-1-1"
	res := diffWaysToCompute(input)
	fmt.Println(res)
	input = "2*3-4*5"
	res = diffWaysToCompute(input)
	fmt.Println(res)
}
