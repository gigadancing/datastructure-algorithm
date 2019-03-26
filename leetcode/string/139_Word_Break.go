package string

// 139.Word Break
// Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, determine if s can be
// segmented into a space-separated sequence of one or more dictionary words.
// Note:
//     The same word in the dictionary may be reused multiple times in the segmentation.
//     You may assume the dictionary does not contain duplicate words.
//
// Example 1:
//     Input: s = "leetcode", wordDict = ["leet", "code"]
//     Output: true
//     Explanation: Return true because "leetcode" can be segmented as "leet code".
//
// Example 2:
//     Input: s = "applepenapple", wordDict = ["apple", "pen"]
//     Output: true
//     Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
//                  Note that you are allowed to reuse a dictionary word.
// Example 3:
//     Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
//     Output: false
func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dict := make(map[string]string, 0)
	for _, w := range wordDict {
		dict[w] = w
	}
	tmp := " " + s
	f := make([]bool, n+1)
	f[0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if f[j] {
				str := tmp[j+1 : j+1+i-j]
				if _, ok := dict[str]; ok {
					f[i] = true
					break
				}
			}
		}
	}

	return f[n]
}

//
func wordBreak2(s string, wordDict []string) bool {
	dict := make(map[string]bool)
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for _, k := range wordDict {
		dict[k] = true
	}

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dict[string(s[j:i])] && dp[j] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}
