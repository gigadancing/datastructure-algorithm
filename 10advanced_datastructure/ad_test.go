package ad

import (
	"fmt"
	"testing"
)

func TestNewTrieNode(t *testing.T) {
	root := TrieNode{}
	n1 := TrieNode{}
	n2 := TrieNode{}
	n3 := TrieNode{}
	root.child['a'-'a'] = &n1
	root.child['b'-'a'] = &n2
	root.child['e'-'a'] = &n3

	n4 := TrieNode{}
	n5 := TrieNode{}
	n6 := TrieNode{}
	n1.child['b'-'a'] = &n4
	n2.child['c'-'a'] = &n5
	n3.child['f'-'a'] = &n6

	n7 := TrieNode{}
	n8 := TrieNode{}
	n9 := TrieNode{}
	n10 := TrieNode{}
	n4.child['c'-'a'] = &n7
	n4.child['d'-'a'] = &n8
	n5.child['d'-'a'] = &n9
	n6.child['g'-'a'] = &n10
	n7.isEnd = true
	n8.isEnd = true
	n9.isEnd = true
	n10.isEnd = true

	n11 := TrieNode{}
	n7.child['d'-'a'] = &n11
	n11.isEnd = true

	preorderTrie(&root, 0)
}

func TestGetAllWord(t *testing.T) {
	root := TrieNode{}
	n1 := TrieNode{}
	n2 := TrieNode{}
	n3 := TrieNode{}
	root.child['a'-'a'] = &n1
	root.child['b'-'a'] = &n2
	root.child['e'-'a'] = &n3
	n2.isEnd = true

	n4 := TrieNode{}
	n5 := TrieNode{}
	n6 := TrieNode{}
	n1.child['b'-'a'] = &n4
	n2.child['c'-'a'] = &n5
	n3.child['f'-'a'] = &n6

	n7 := TrieNode{}
	n8 := TrieNode{}
	n9 := TrieNode{}
	n10 := TrieNode{}
	n4.child['c'-'a'] = &n7
	n4.child['d'-'a'] = &n8
	n5.child['d'-'a'] = &n9
	n6.child['g'-'a'] = &n10
	n7.isEnd = true
	n8.isEnd = true
	n9.isEnd = true
	n10.isEnd = true

	n11 := TrieNode{}
	n7.child['d'-'a'] = &n11
	n11.isEnd = true

	word := make([]byte, 0)
	wordList := make([]string, 0)
	GetAllWord(&root, &word, &wordList)
	fmt.Println(wordList)
}

func TestTrieTree(t *testing.T) {
	trie := NewTrieTree()
	trie.Insert("abc")
	trie.Insert("abcd")
	trie.Insert("abd")
	trie.Insert("b")
	trie.Insert("acd")
	trie.Insert("efg")
	word := make([]byte, 0)
	wordList := make([]string, 0)
	GetAllWord(trie.root, &word, &wordList)
	fmt.Println(wordList)

	fmt.Println(trie.Search("xxxx"))
	fmt.Println(trie.Search("abcd"))
	fmt.Println(trie.Search("b"))
	fmt.Println(trie.Search("abc"))
}

func TestWordDictionary(t *testing.T) {
	wd := NewWordDictionary()
	wd.addWord("bad")
	wd.addWord("dad")
	wd.addWord("mad")
	fmt.Println(wd.search("pad"))
	fmt.Println(wd.search("bad"))
	fmt.Println(wd.search(".ad"))
	fmt.Println(wd.search("b.."))
}

func TestFindCircleNum(t *testing.T) {
	m := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	fmt.Println(FindCircleNum(m))
	m = [][]int{{1, 1, 0}, {1, 1, 1}, {0, 1, 1}}
	fmt.Println(FindCircleNum(m))
}

func TestBuildSegmentTree(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5}
	values := make([]int, 20)
	buildSegmentTree(&values, nums, 0, 0, len(nums)-1)
	fmt.Println(values)
	printSegmentTree(values, 0, 0, len(nums)-1, 0)
}
