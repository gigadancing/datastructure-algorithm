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

// 例2. 打家劫舍
// 在一条直线上，有n个房屋，每个房屋中有数量不等的财宝，有一个盗贼希望从房屋中盗取财宝，由于房屋中有报警器，如果同时从两个相邻的房屋中盗
// 取财宝就会触发报警器。问在不处罚报警器的前提下，最多可以获取多少财宝。
// 例如：
// [5,2,6,3,1,7]，最多为5+6+7=18
func Rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(nums)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 例3. 最大子段和
// 给定一个数组，求这个数组的连续子数组中，最大的那一段的和。
// 例如：
// [-2,1,-3,4,-1,2,1,-5,4]，连续子数组如：
// [-2,1]、[1,-3,4,-1]、[4,-1,2,1]，...，最大的是[4,-1,2,1]，和为6
// 思路：
// 求n个数的数组的最大子段和，转换为分别求以第1个、第2个...第i个...第n个数字结尾的最大子段和，再找出n个结果中最大的即为结果。
func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = 0
	maxSum := dp[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		if maxSum < dp[i] {
			maxSum = dp[i]
		}
	}

	return maxSum
}

// 例4. 找零钱
// 已知不同面值的钞票，如何用最少数量的钞票组成某个金额，求使用钞票的最少数量。如果任意数量的已知面额钞票都无法组成该金额，则返回-1。
// 例如：
// 钞票面值：[1,2,5]，金额11=5+5+1，需要3张
// 钞票面值：[2]，金额3，无法组成，返回-1
// 钞票面值：[1,2,5,7,10]，金额14=7+7，需要2张
// 贪心思想（每次使用最大的面值）可否？
// 例如：
// [1,2,5,10]，14，10+2+2=14，ok
// [1,2,5,7,10],14,10+2+2=14, error，（最优解为7+7=14）
// 因此，贪心思想不可行。
// 思路：
// dp[i]代表金额i的最优解（组成金额i的最少钞票数量）
// 数组dp存储从金额1到金额M的最优解
// 由于金额i可由：
// 金额i-coins[0]与coins[0]
// 金额i-coins[1]与coins[1]
// 金额i-coins[2]与coins[2]
// 金额i-coins[3]与coins[3]
// ...
// 故状态i可由状态i-coins[0]、i-coins[1]、i-coins[2]、i-coins[3]...转换得到
// 即dp[i] = min(dp[i-coins[0]], dp[i-coins[1]], dp[i-coins[2]], dp[i-coins[3]], ...) + 1
func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = -1 // 初始化所有的金额的最优解为-1
	}
	dp[0] = 0 // 金额0的最优解为0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ { // 循环各个面值，找到dp[i]的最优解
			if i-coins[j] >= 0 && dp[i-coins[j]] != -1 { // 递推条件
				if dp[i] == -1 || dp[i] > dp[i-coins[j]] {
					dp[i] = dp[i-coins[j]] + 1 // 递推公式
				}
			}
		}
	}
	return dp[amount]
}

// 例5. 三角形
// 给定一个二维数组，其保存了一个数字三角想，求从数字三角形顶端到底端最小的路径之和，每次可以向下走相邻的两个位置。
// [2]
// [3 4]
// [6 5 7]
// [4 1 8 3]
func MinimumTotal(triangle [][]int) int {
	// 初始化和triangle大小相同的数组
	dp := make([][]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		arr := make([]int, len(triangle[i]))
		dp[i] = arr
	}

	for i := len(triangle) - 1; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			if i+1 < len(triangle) && j+1 < len(triangle[i+1]) {
				dp[i][j] += min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
			} else {
				dp[i][j] = triangle[i][j]
			}
		}
	}

	return dp[0][0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 例6. 最长上升子序列
// 已知一个未排序数组，求这个数组的最长上升子序列的长度。
// 例如：
// [1,3,2,3,1,4]，其中有很多上升子序列，如[1,3]、[1,2,3]、[1,2,3,4]，最长上升子序列的长度为4。分别考虑O(2^n)和O(nlogn)两种复杂度
// 算法。
// 思路：
// 若dp[i]代表以第i个元素结尾的最长上升子序列长度，dp[i-1]代表以第i-1个元素结尾的最长上升子序列长度，那么nums[i]一定是dp[i]中所有的
// 最长上升子序列中最大的元素（因为在末尾）
// dp[0]=1,[1]
// dp[1]=2,[1,3]
// dp[2]=2,[1,2]
// dp[3]=3,[1,2,3]
// dp[4]=1,[1]
// dp[5]=4,[1,2,3,4]
// 最终结果为dp[0],dp[1],dp[2],...,dp[i],...,dp[n-1]中的最大值（与最大子段和相似之处）。
// 算法：
// 设置动态规划数组dp，第i个元素代表以第i个元素结尾的最长上升子序列的长度。动态规划边界为dp[0]=1
// 初始化最长上升子序列的长度LIS=1
// 从1到n-1，循环i，计算dp[i]:
//     从0至i-1，循环j，若nums[i]>nums[j]，说明nums[i]可以放置在nums[j]的后面，组成最长上升子序列:
//         若dp[i]<dp[j]+1:
//             dp[i]=dp[j]+1
func LengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	maxLen := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1                // 每个元素的最长子序列至少是自己本身，即长度为1
		for j := 0; j < i; j++ { // 遍历0到i-1
			if nums[i] > nums[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
			if maxLen < dp[i] {
				maxLen = dp[i]
			}
		}
	}

	return maxLen
}

// 设置一个栈，stack[i]代表长度为i+1的上升子序列最后一个元素的最小可能的取值，即若要组成一个i+2长度的上升子序列，需要一个大于stack[i]
// 的元素。最终栈的大小即为最长上升子序列的长度。
// 算法：
// 1. 设置一个栈，将nums[0]入栈
// 2. 从1到n-1遍历数组：
//    若nums[i]>栈顶：将nums[i]入栈
//    否则：从栈底遍历至栈顶，若栈元素大于等于nums[i]：用nums[i]替换该元素
// 3. 返回栈中元素个数
func LengthOfLIS2(nums []int) int {
	s := make([]int, len(nums))
	s[0] = nums[0]
	pos := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > s[pos] {
			s[pos+1] = nums[i]
			pos++
		} else {
			for j := 0; j <= pos; j++ {
				if s[j] >= nums[i] {
					s[j] = nums[i]
					break
				}
			}
		}
	}

	return pos + 1 // pos是从0开始的，故元素个数为pos+1
}

// 例7. 最小路径和
// 已知一个二维数组，里面存储了非负整数，找到从左上角到右下角的一条路径，使得路径上的和最小（移动过程中只能向下或向右）。
// 1  3  1
// 1  5  1
// 4  2  1
// 故最短路径为[1,3,1,1,1]，和为7。
func MinPathSum(grid [][]int) int {
	rows := len(grid)           // 行数
	cols := len(grid[0])        // 列数
	dp := make([][]int, rows)   // 二维数组
	for i := 0; i < rows; i++ { // 初始化二维数组dp
		dp[i] = make([]int, cols)
	}
	dp[0][0] = grid[0][0]

	for i := 1; i < cols; i++ { // 第一行中的位置，只能从左边到达，先将求出dp中第一行
		dp[0][i] = grid[0][i] + dp[0][i-1]
	}
	for i := 1; i < rows; i++ { // 第一列中的位置，只能从上边到达，先将求出dp中第一列
		dp[i][0] = grid[i][0] + dp[i-1][0]
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			dp[i][j] = grid[i][j] + min(dp[i][j-1], dp[i-1][j])
		}
	}

	return dp[rows-1][cols-1]
}

// 例8. 地牢游戏
// 已知一个二维数组，左上角代表骑士位置，右下角代表公主位置。二维数组中存储整数，正数给骑士增加生命值，负数给骑士减少生命值，问骑士至少
// 是多少生命值，才可保证骑士在行走的过程中至少保持生命值为1。（骑士只能向下或向右行走）
// -2  -3   3
// -5  -10  1
// 10  30  -5
// 从右小角向左上角递推，dp[i][j]代表从上一个位置(i,j-1)或(i-1,j)走到该位置骑士保持生命值为1，需要的最少生命值
func CalculateMinimumHP(dungeon [][]int) int {
	rows := len(dungeon)
	cols := len(dungeon[0])
	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, cols)
	}
	dp[rows-1][cols-1] = max(1, 1-dungeon[rows-1][cols-1])

	// 递推最后一行
	for col := cols - 2; col >= 0; col-- {
		dp[rows-1][col] = max(1, dp[rows-1][col+1]-dungeon[rows-1][col])
	}

	// 递推最后一列
	for row := rows - 2; row >= 0; row-- {
		dp[row][cols-1] = max(1, dp[row+1][cols-1]-dungeon[row][cols-1])
	}

	for row := rows - 2; row >= 0; row-- {
		for col := cols - 2; col >= 0; col-- {
			below := max(1, dp[row+1][col]-dungeon[row][col]) // 从下面推
			right := max(1, dp[row][col+1]-dungeon[row][col]) // 从右面推
			dp[row][col] = min(below, right)
		}
	}

	return dp[0][0]
}
