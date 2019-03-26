package string

import (
	"strings"
)

func wordBreak(s string, wordDict []string) bool {
	arr1 := getArray(s)
	arr2 := getArray(strings.Join(wordDict, ""))

	for i := 0; i < 26; i++ {
		if arr1[i] < arr2[i] {
			return false
		}

		if (arr1[i] > arr2[i]) && arr2[i] == 0 {
			return false
		}
	}

	return true
}

func getArray(s string) [26]byte {
	arr := [26]byte{}
	for _, ch := range s {
		if ch <= 90 {
			ch += 32
		}
		arr[ch-'a']++
	}
	return arr
}
