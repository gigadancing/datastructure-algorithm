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
			value := make([]string, 0)
			value = append(value, word)
			anagram[string(w)] = value
		}
	}
	for _, v := range anagram {
		res = append(res, v)
	}
	return res
}

func generateKey(str string) [26]byte {
	key := [26]byte{}
	for _, ch := range str {
		key[ch-'a'] = 1
	}
	return key
}

func GroupAnagram2(words []string) [][]string {
	res := make([][]string, 0)
	anagram := make(map[[26]byte][]string, 0)

	for _, word := range words {
		key := generateKey(word)
		if _, ok := anagram[key]; ok {
			anagram[key] = append(anagram[key], word)
		} else {
			value := make([]string, 0)
			value = append(value, word)
			anagram[key] = value
		}
	}
	for _, v := range anagram {
		res = append(res, v)
	}
	return res
}

// 例4. 无重复字符的最长字串
// 已知一个字符串，求用该字符串的无重复字符的最长子串的长度。
// 例如：
// s="abcabcbb"->"abc",3
// s="bbbbb"->"b",1
// s="pwwkew"->"wke",3 注意"pwke"是子序列而非字串
func LengthOfLongestSubstring(str string) (int, string) {
	maxLen := 0
	tmpStr := ""
	maxSubstr := ""
	charMap := [128]byte{}
	begin := 0
	for i, ch := range str {
		charMap[ch]++
		if charMap[ch] == 1 { // 字符只出现了一次，说明未重复
			tmpStr += string(ch)
			if maxLen < len(tmpStr) {
				maxLen = len(tmpStr)
				maxSubstr = tmpStr
			}
		} else {
			for begin < i && charMap[ch] > 1 {
				charMap[str[begin]]--
				begin++
			}
			tmpStr = string([]byte(str)[begin : i+1]) // 更新tmpStr
		}
	}

	return maxLen, maxSubstr
}

// 例5. 重复DNA序列
// 将DNA序列看作是只包含['A','G','C','T']4个字符的字符串，给一个DNA字符串，找到所有长度为10的且出现超过一次的子串。
// 例如：
// s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
// return ["AAAAACCCCC","CCCCCAAAAA"]
// s = "AAAAAAAAAAAA"
// return ["AAAAAAAAAA"]
func FindRepeatedDnaSequences(str string) []string {
	res := make([]string, 0)
	substrMap := make(map[string]int, 0)
	for i := 0; i < len(str) && i+10 < len(str); i++ {
		substr := str[i : i+10]
		if _, ok := substrMap[substr]; ok {
			substrMap[substr]++
		} else {
			substrMap[substr] = 1
		}
	}

	for k, v := range substrMap {
		if v > 1 {
			res = append(res, k)
		}
	}

	return res
}

// 例6. 最小窗口子串
// 已知字符串S与字符串T，求在S中最小窗口（区间），使得这个区间中包含字符串T中的所有字符。
// 例如：
// s="ADOBECODEBANC";T="ABC"
// 包含T的子区间中，有"ADOBEC","CODEBA","BANC"等，最小窗口区间是"BANC"
func MinWindow(s, t string) string {
	minStr := "" // 最小窗口字符串
	begin := 0   // 最小窗口起始位置
	current, target := [128]byte{}, [128]byte{}
	m := make([]byte, 0)

	for _, ch := range t { // 遍历t建立映射关系target
		target[ch]++
	}
	for v := range target { // 遍历target，将其中出现的字符加入到数组m中
		if v > 0 {
			m = append(m, byte(v))
		}
	}

	for i, ch := range s {
		current[ch]++

		for begin < i {
			c := s[begin]
			if target[c] == 0 {
				begin++
			} else if current[c] > target[c] {
				current[c]--
				begin++
			} else {
				break
			}
		}

		if isWindowOk(current, target, m) { // 检查此时窗口是否包含target
			newWindowLen := i - begin + 1                   // 当前字符串长度
			if minStr == "" || len(minStr) > newWindowLen { // 结果为空字符串或当前字符串更短时，更新字符串
				minStr = s[begin : i+1]
			}
		}
	}

	return minStr
}

// 检查s是否包含t
func isWindowOk(s, t [128]byte, m []byte) bool {
	for _, b := range m {
		if s[b] < t[b] {
			return false
		}
	}
	return true
}
