package ad

import (
	"fmt"
)

const N = 26

type TrieNode struct {
	child [N]*TrieNode
	isEnd bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		isEnd: false,
	}
}

func preorderTrie(node *TrieNode, layer int) {
	if node == nil {
		return
	}
	for i, ch := range node.child {
		if ch != nil {
			for j := 0; j < layer; j++ {
				fmt.Printf("---")
			}
			fmt.Printf("%c", 'a'+i)
			if ch.isEnd {
				fmt.Printf("(end)")
			}
			fmt.Println()
			preorderTrie(ch, layer+1)
		}
	}
}

// Trie树获取全部单词
// 深度搜索trie树，对于正在搜索的节点node：
// 遍历该节点的26个孩子指针child[i]
//     如果指针不为空：将该child[i]对应的字符入栈
//     如果该孩子的isEnd为真，说明这个位置是单词结尾，从栈底到栈顶对栈进行遍历，生成字符串，将它保存到结果数组中
// 深度搜索child[i]
// 弹出栈顶字符
func GetAllWord(node *TrieNode, word *[]byte, wordList *[]string) {
	if node == nil {
		return
	}
	for i, ch := range node.child {
		if ch != nil {
			*word = append(*word, byte('a'+i)) // 字符入栈
			if ch.isEnd {
				*wordList = append(*wordList, string(*word)) // 单词加入结果数组
			}
			GetAllWord(ch, word, wordList) // 继续深搜
			*word = (*word)[:len(*word)-1] // 弹出栈顶字符
		}
	}
}

// 字典树
type TrieTree struct {
	root *TrieNode
}

//
func NewTrieTree() *TrieTree {
	return &TrieTree{
		root: NewTrieNode(),
	}
}

// 插入
func (t *TrieTree) Insert(word string) {
	node := t.root
	for _, char := range word {
		pos := byte(char) - 'a'
		if node.child[pos] == nil {
			node.child[pos] = NewTrieNode()
		}
		node = node.child[pos]
	}
	node.isEnd = true
}

// 查找
func (t *TrieTree) Search(word string) bool {
	node := t.root
	for _, char := range word {
		pos := byte(char) - 'a'
		if node.child[pos] == nil {
			return false
		}
		node = node.child[pos]
	}
	return node.isEnd
}

// 以某个前缀开始
func (t *TrieTree) StartsWith(prefix string) bool {
	node := t.root
	for _, char := range prefix {
		pos := byte(char) - 'a'
		if node.child[pos] == nil {
			return false
		}
		node = node.child[pos]
	}
	return true
}

// 例2. 添加与查找单词
// 设计一个数据结构，支持如下两种操作：
// 1）添加单词，addWord(word)
// 2）查找单词，bool search(word)
// 添加单词只包含小写字符'a'-'z'
// 搜索单词时，可以按照普通的方式搜索单词（原始单词）或正则表达式方式搜索单词；搜索单词时只包含小写字符'a'-'z'或'.'，'.'代表任意一个
// 小写字符
// 例如：
// addWord("bad")
// addWord("dad")
// addWord("mad")
// search("pad") -> false
// search("bad") -> true
// search(".ad") -> true
// search("b..") -> true
type WordDictionary struct {
	tree *TrieTree
}

// 构造函数
func NewWordDictionary() *WordDictionary {
	return &WordDictionary{
		tree: NewTrieTree(),
	}
}

// 添加单词
func (wd *WordDictionary) addWord(word string) {
	wd.tree.Insert(word)
}

// 搜索单词
func (wd *WordDictionary) search(word string) bool {
	return searchTrie(wd.tree.root, word, 0)
}

// 回溯深搜
func searchTrie(node *TrieNode, word string, pos int) bool {
	if pos == len(word) {
		if node.isEnd {
			return true
		}
		return false
	}

	if word[pos] == '.' {
		for i := 0; i < N; i++ {
			if node.child[i] != nil && searchTrie(node.child[i], word, pos+1) {
				return true
			}
		}
	} else {
		index := word[pos] - 'a'
		if node.child[index] != nil && searchTrie(node.child[index], word, pos+1) {
			return true
		}
	}

	return false
}

// 例3. 朋友圈
// 有N个同学，他们之间有些是朋友，有些不是。“友谊”是可以传递的，例如A与B是朋友，B与C是朋友，那么A与C也是朋友；朋友圈就是完成“友谊”传递
// 后的一组朋友。给定N*N的矩阵代表同学间是否是朋友，如果M[i][j]=1代表第i个学生与第j个学生是朋友，否则不是。求朋友圈的个数。
// Input:
// [[1,1,0],
//  [1,1,0],
//  [0,0,1]]
// Output:2
// Input:
// [[1,1,0],
//  [1,1,1],
//  [0,1,1]]
// Output:1
func FindCircleNum(m [][]int) int {
	ds := NewDisjoinSet(len(m))
	for i := 0; i < len(m); i++ {
		for j := i + 1; j < len(m); j++ {
			if m[i][j] != 0 {
				ds.union(i, j)
			}
		}
	}
	return ds.count
}

// 数组实现并查集，复杂度为O(n)
type DisjoinSet2 struct {
	id []int
}

// 构造函数
func NewDisjoinSet2(n int) *DisjoinSet2 {
	id := make([]int, 0)
	for i := 0; i < n; i++ { // 每个元素单独构成一个集合，编号i的元素属于集合i
		id = append(id, i)
	}
	return &DisjoinSet2{
		id: id,
	}
}

// 查询元素p属于哪个集合
func (ds *DisjoinSet2) find(p int) int {
	return ds.id[p]
}

// 合并元素p和q
func (ds *DisjoinSet2) union(p, q int) {
	pid := ds.find(p)
	qid := ds.find(q)
	if pid == qid { // 两元素属于同一集合
		return
	}
	for i := 0; i < len(ds.id); i++ {
		if ds.id[i] == pid { // 讲所有属于pid的元素改为属于qid
			ds.id[i] = qid
		}
	}
}

// 优化后的并查集
type DisjoinSet struct {
	count int
	id    []int
	size  []int
}

// 构造函数
func NewDisjoinSet(n int) *DisjoinSet {
	id := make([]int, 0)
	size := make([]int, 0)
	for i := 0; i < n; i++ {
		id = append(id, i)
		size = append(size, 1)
	}

	return &DisjoinSet{
		id:    id,
		size:  size,
		count: n,
	}
}

// 查询
func (ds *DisjoinSet) find(p int) int {
	for p != ds.id[p] {
		ds.id[p] = ds.id[ds.id[p]]
		p = ds.id[p]
	}
	return p
}

// 合并
func (ds *DisjoinSet) union(p, q int) {
	i := ds.find(p)
	j := ds.find(q)
	if i == j {
		return
	}
	if ds.size[i] < ds.size[j] {
		ds.id[i] = j
		ds.size[j] += ds.size[i]
	} else {
		ds.id[j] = i
		ds.size[i] += ds.size[j]
	}
	ds.count--
}
