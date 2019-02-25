package ad

import "fmt"

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
