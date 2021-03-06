package hts

import (
	"fmt"
	"testing"
)

func TestHashTable(t *testing.T) {
	tableLen := 11 // 长度为质数
	nodes := make([]*ListNode, 0)
	testNums := []int{1, 1, 4, 9, 20, 30, 150, 500}
	hashTable := make([]*ListNode, tableLen)
	for _, num := range testNums {
		nodes = append(nodes, NewListNode(num))
	}

	for _, node := range nodes {
		Insert(hashTable, node, tableLen)
	}

	for i, node := range hashTable {
		fmt.Printf("%3d : ", i)
		for p := node; p != nil; p = p.Next {
			fmt.Printf("%3d", p.Val)
		}
		fmt.Println()
	}

	for i := 0; i < 10; i++ {
		exist := Search(hashTable, tableLen, i)
		fmt.Printf("%3d : %v\n", i, exist)
	}

}

func TestLongestPalindrome(t *testing.T) {
	s := "abccccddaawwwww"
	n := LongestPalindrome(s)
	fmt.Println(n)
	s = "aaabbb"
	n = LongestPalindrome(s)
	fmt.Println(n)
	s = "iijj"
	n = LongestPalindrome(s)
	fmt.Println(n)
}

func TestWordPattern(t *testing.T) {
	pattern := "abba"
	str := "dog cat cat dog"
	matched := WordPattern(pattern, str)
	fmt.Println(matched)
	pattern = "abba"
	str = "dog cat cat fish"
	matched = WordPattern(pattern, str)
	fmt.Println(matched)
	pattern = "abcd"
	str = "dog cat cat fish"
	matched = WordPattern(pattern, str)
	fmt.Println(matched)
	pattern = "abbc"
	str = "dog cat cat fish"
	matched = WordPattern(pattern, str)
	fmt.Println(matched)
	pattern = "ab"
	str = "dog cat cat fish"
	matched = WordPattern(pattern, str)
	fmt.Println(matched)

}

func TestGroupAnagram(t *testing.T) {
	words := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := GroupAnagram(words)
	for _, arr := range res {
		for _, w := range arr {
			fmt.Printf("%v ", w)
		}
		fmt.Printf("\n")
	}
	fmt.Println("-----------------------------")
	res = GroupAnagram2(words)
	for _, arr := range res {
		for _, w := range arr {
			fmt.Printf("%v ", w)
		}
		fmt.Printf("\n")
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "abcabcbb"
	maxLen, substr := LengthOfLongestSubstring(s)
	fmt.Println(maxLen, substr)
	s = "bbbbb"
	maxLen, substr = LengthOfLongestSubstring(s)
	fmt.Println(maxLen, substr)
	s = "pwwkew"
	maxLen, substr = LengthOfLongestSubstring(s)
	fmt.Println(maxLen, substr)
}

func TestFindRepeatedDnaSequences(t *testing.T) {
	s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	res := FindRepeatedDnaSequences(s)
	fmt.Println(res)
	s = "AAAAAAAAAAAA"
	res = FindRepeatedDnaSequences(s)
	fmt.Println(res)
}

func TestMinWindow(t *testing.T) {
	s := "ADOBECODEBANC"
	tt := "ABC"
	minStr := MinWindow(s, tt)
	fmt.Println(minStr)
	tt = "ABCD"
	minStr = MinWindow(s, tt)
	fmt.Println(minStr)
}
