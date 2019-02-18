package dp

// 例1. 爬楼梯
// 在爬楼梯时，每次可向上走1个台阶或2个台阶，问有n阶楼梯时，有多少种上楼方式。
func ClimbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	return ClimbStairs(n-1) + ClimbStairs(n-2) // 递归会超时
}

// 优化（动态规划）
// 1. 设置递归数组dp[0...n]，dp[i]代表达到第i阶有多少种走法，初始化数组元素为0。
// 3. 设置达到第1阶有1种走法，到达第2阶有2种走法
// 3. 利用i循环递推从第3阶到第n阶结果
//    到达第i阶的走法 = 达到第i-1阶的走法 + 达到第i-2阶的走法
func ClimbStairsDP(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}
