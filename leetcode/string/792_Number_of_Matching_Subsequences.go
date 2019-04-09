package string

// 792.Number of Matching Subsequences
// Given string S and a dictionary of words words, find the number of words[i] that is a subsequence of S.
//
// Example :
// Input:
//       S = "abcde"
//       words = ["a", "bb", "acd", "ace"]
// Output: 3
// Explanation: There are three words in words that are a subsequence of S: "a", "acd", "ace".
// Note:
//      All words in words and S will only consists of lowercase letters.
//      The length of S will be in the range of [1, 50000].
//      The length of words will be in the range of [1, 5000].
//      The length of words[i] will be in the range of [1, 50].
func numMatchingSubseq(S string, words []string) int {
	count := 0
	for _, w := range words {
		if isMatch(S, w) {
			count++
		}
	}
	return count
}

func isMatch(S, word string) bool {
	start := 0
	for _, ch := range word {
		found := false
		for i := start; i < len(S); i++ {
			if int32(S[i]) == ch {
				found = true
				start = i + 1
				break
			}
			if !found {
				return false
			}
		}
	}
	return true
}
