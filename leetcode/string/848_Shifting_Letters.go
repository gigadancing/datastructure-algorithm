package string

// 848.Shifting Letters
// We have a string S of lowercase letters, and an integer array shifts.
// Call the shift of a letter, the next letter in the alphabet, (wrapping around so that 'z' becomes 'a').
// For example, shift('a') = 'b', shift('t') = 'u', and shift('z') = 'a'.
// Now for each shifts[i] = x, we want to shift the first i+1 letters of S, x times.
// Return the final string after all such shifts to S are applied.
//
// Example 1:
//
// Input: S = "abc", shifts = [3,5,9]
// Output: "rpl"
// Explanation:
// We start with "abc".
// After shifting the first 1 letters of S by 3, we have "dbc".
// After shifting the first 2 letters of S by 5, we have "igc".
// After shifting the first 3 letters of S by 9, we have "rpl", the answer.
// Note:
// 1 <= S.length = shifts.length <= 20000
// 0 <= shifts[i] <= 10 ^ 9
func shiftingLetters(S string, shifts []int) string {
	bytes := []byte(S)
	sum := 0
	for i := len(shifts) - 1; i >= 0; i-- {
		shifts[i] = shifts[i] % 26
		sum += shifts[i]
		bytes[i] = 'a' + byte((int(bytes[i])-'a'+sum)%26)
	}

	return string(bytes)
}

// leetcode牛逼代码
func shiftingLetters2(S string, shifts []int) string {
	n := len(S)
	buf := make([]byte, n)
	sum := 0
	for i := n - 1; i >= 0; i-- {
		sum += shifts[i] % 26
		sum %= 26
		shifts[i] = sum
	}
	for i, c := range S {
		buf[i] = byte(int('a') + (int(c)+shifts[i]-int('a'))%26)
	}
	return string(buf)
}
