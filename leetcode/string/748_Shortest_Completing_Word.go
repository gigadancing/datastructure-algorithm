package string

import (
	"github.com/ethereum/go-ethereum/common/math"
	"strings"
)

// 748.Shortest Completing Word
// Find the minimum length word from a given dictionary words, which has all the letters from the string licensePlate.
// Such a word is said to complete the given string licensePlate
//
// Here, for letters we ignore case. For example, "P" on the licensePlate still matches "p" on the word.
//
// It is guaranteed an answer exists. If there are multiple answers, return the one that occurs first in the array.
//
// The license plate might have the same letter occurring multiple times. For example, given a licensePlate of "PP",
// the word "pair" does not complete the licensePlate, but the word "supper" does.
//
// Example 1:
// Input: licensePlate = "1s3 PSt", words = ["step", "steps", "stripe", "stepple"]
// Output: "steps"
// Explanation:
//     The smallest length word that contains the letters "S", "P", "S", and "T".
//     Note that the answer is not "step", because the letter "s" must occur in the word twice.
//     Also note that we ignored case for the purposes of comparing whether a letter exists in the word.
//
// Example 2:
// Input: licensePlate = "1s3 456", words = ["looks", "pest", "stew", "show"]
// Output: "pest"
// Explanation:
//     There are 3 smallest length words that contains the letters "s".
//     We return the one that occurred first.
// Note:
//     licensePlate will be a string with length in range [1, 7].
//     licensePlate will contain digits, spaces, or letters (uppercase or lowercase).
//     words will have a length in the range [10, 1000].
//     Every words[i] will consist of lowercase letters, and have length in range [1, 15].
// 首先想到的解法：先求出licensePlate中每个字符出现的次数，再和words中的一一对比，看是否满足题目条件。
// 提交通过后，但是效率不高。
func shortestCompletingWord(licensePlate string, words []string) string {
	m := make(map[int32]int)
	for _, ch := range licensePlate { // 统计每个字符出现的次数
		if ch >= 65 && ch <= 90 {
			m[ch+32]++
		} else if ch >= 97 && ch <= 122 {
			m[ch]++
		}
	}

	res := make([]string, 0)
	for _, w := range words {
		if isContained(m, w) {
			res = append(res, w)
		}
	}

	if len(res) == 0 {
		return ""
	}
	minStr := res[0]
	minLen := len(res[0])
	for _, s := range res {
		if len(s) < minLen {
			minStr = s
			minLen = len(s)
		}
	}

	return minStr
}

// m中的字符是否满足被包含在w中
func isContained(m map[int32]int, w string) bool {
	chars := make(map[int32]int)
	for _, ch := range w {
		chars[ch]++
	}

	for k, v := range m {
		if _, ok := chars[k]; !ok {
			return false
		} else if v > chars[k] {
			return false
		}
	}

	return true
}

// 第二种解法
// 不用map，改用固定长度的数组
func shortestCompletingWord2(licensePlate string, words []string) string {
	lp := [26]int32{}
	for _, ch := range licensePlate {
		if isAlpha(ch) {
			lp[toLower(ch)-'a']++
		}
	}
	ans := ""
	minLen := math.MaxInt32
	for _, w := range words {
		if len(w) >= minLen {
			continue
		}
		if !match(lp, w) {
			continue
		}
		ans = w
		minLen = len(w)
	}

	return ans
}

//
func match(lp [26]int32, w string) bool {
	bytes := [26]int32{}
	for _, ch := range w {
		bytes[ch-'a']++
	}
	for i := 0; i < 26; i++ {
		if lp[i] > bytes[i] {
			return false
		}
	}
	return true
}

// 判断字符是否是字母
func isAlpha(c int32) bool {
	if (c >= 65 && c <= 90) || (c >= 97 && c <= 122) {
		return true
	}
	return false
}

// 将字符转换成小写
func toLower(c int32) int32 {
	if c >= 65 && c <= 90 {
		return c + 32
	}
	return c
}

// leetcode该题牛逼代码
func shortestCompletingWord3(licensePlate string, words []string) string {
	freq := make([]int, 26)
	licensePlate = strings.ToLower(licensePlate)
	for _, c := range licensePlate {
		if 'a' <= c && c <= 'z' {
			freq[c-'a']++
		}
	}

	res := ""
	for _, s := range words {
		temp := make([]int, 26)
		copy(temp, freq)
		if (len(res) == 0 || len(s) < len(res)) && check(temp, s) {
			res = s
		}
	}

	return res
}

func check(freq []int, s string) bool {
	for _, c := range s {
		freq[c-'a']--
	}

	for _, f := range freq {
		if f > 0 {
			return false
		}
	}

	return true
}
