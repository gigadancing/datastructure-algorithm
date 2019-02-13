package hts

import (
	"sort"
	"strings"
)

// 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// 链表节点构造函数
func NewListNode(v int) *ListNode {
	return &ListNode{
		Val: v,
	}
}

// 哈希函数
func HashFunc(key, tableLen int) int {
	return key % tableLen
}

// 插入节点
func Insert(hashTable []*ListNode, node *ListNode, tableLen int) {
	index := HashFunc(node.Val, tableLen)
	node.Next = hashTable[index]
	hashTable[index] = node
}

// 查找值
func Search(hashTable []*ListNode, tableLen, value int) bool {
	index := HashFunc(value, tableLen)
	if hashTable[index] == nil {
		return false
	}
	for p := hashTable[index]; p != nil; p = p.Next {
		if p.Val == value {
			return true
		}
	}
	return false
}

// 哈希表的表长一般为质数，使取余的结果非常分散，不易重复

// 例1. 最长回文串
// 已知一个只包含大小写字符的字符串，求用该字符串中的字符可以生成的最长回文字符串长度。
// 例如：s="abccccddaa"，可以生成的最长回文串长度为9，如"dccaaaccd"、"adccbccda"、"acdcacdca"等都是正确的。
func LongestPalindrome(s string) int {
	bytes := [128]byte{0}
	maxLength := 0 // 回文字符串中偶数部分
	flag := 0      // 是否有中心点
	for _, v := range s {
		bytes[v-'0']++
	}
	for i := 0; i < 128; i++ {
		if bytes[i]%2 == 0 {
			maxLength += int(bytes[i])
		} else {
			maxLength += int(bytes[i]) - 1
			flag = 1
		}
	}

	return maxLength + flag
}

// 例2. 词语模式
// 已知字符串pattern与字符串，确忍str是否与pattern匹配。str与pattern匹配代表字符串str中的单词与pattern中的字符一一对应。
// （其中pattern只包含小写字符，str中的单词只包含小写字符，使用空格分隔。）
// 例如：
// pattern="abba", str="dog cat cat dog", 匹配
// pattern="abba", str="dog cat cat fish", 不匹配
// pattern="aaaa", str="dog cat cat dog", 不匹配
// pattern="abba", str="dog dog dog dog", 不匹配
// 分析：
// 1. 当拆解出一个单词时，若该单词已出现，则当前单词对应的pattern字符必为该单词曾经对应的pattern字符。
// 2. 当拆解出一个单词时，若该单词未曾出现。
// 3. 单词的个数与pattern字符串中的字符数量相同。
func WordPattern(pattern, str string) bool {
	words := strings.Split(str, " ") // 拆分单词字符串
	if len(pattern) != len(words) {  // pattern长度和单词的个数不相等，不匹配
		return false
	}
	used := [128]byte{}
	wordMap := make(map[string]byte, 0)
	pos := 0 // pattern中字符位置
	for _, w := range words {
		if _, ok := wordMap[w]; ok { // 单词在映射中
			if wordMap[w] != pattern[pos] { // 已有的映射关系无法与当前pattern字符对应
				return false
			}
		} else { // 单词未在映射中
			if used[pattern[pos]] == 1 { // 当前pattern字符已使用
				return false
			}
			wordMap[w] = pattern[pos]
			used[pattern[pos]] = 1
		}
		pos++
	}

	return true
}

type WORD []byte

func (w *WORD) Len() int {
	return len(*w)
}

func (w *WORD) Less(i, j int) bool {
	return (*w)[i] < (*w)[j]
}

func (w *WORD) Swap(i, j int) {
	(*w)[i], (*w)[j] = (*w)[j], (*w)[i]
}

// 例3. 同字符词语分组
// 已知一组字符串，将所有anagram（由颠倒字母顺序而构成的字）放到一起输出。
// 例如：["eat","tea","tan","ate","nat","bat"]
// 返回：[["eat","tea","ate"], ["tan","nat"], ["bat"]]
// 即：字符串里的字符相同，就该分到一组
// 思考：
// 如何建立哈希表，怎样设计哈希表的key和value，就可将字符内容相同的字符串映射到一起？
func GroupAnagram(words []string) [][]string {
	res := make([][]string, 0)
	anagram := make(map[string][]string, 0)
	for _, word := range words {
		w := WORD(word)
		sort.Sort(&w)
		if _, ok := anagram[string(w)]; ok {
			anagram[string(w)] = append(anagram[string(w)], word)
		} else {
			arr := make([]string, 0)
			arr = append(arr, word)
			anagram[string(w)] = arr
		}
	}
	for _, v := range anagram {
		res = append(res, v)
	}
	return res
}
