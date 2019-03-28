package string

import "strconv"

// 241.Different Ways to Add Parentheses
// Given a string of numbers and operators, return all possible results from computing all the different possible ways
// to group numbers and operators. The valid operators are +, - and *.
//
// Example 1:
//     Input: "2-1-1"
//     Output: [0, 2]
//     Explanation:
//         ((2-1)-1) = 0
//         (2-(1-1)) = 2
//
// Example 2:
//     Input: "2*3-4*5"
//     Output: [-34, -14, -10, -10, 10]
//     Explanation:
//         (2*(3-(4*5))) = -34
//         ((2*3)-(4*5)) = -14
//         ((2*(3-4))*5) = -10
//         (2*((3-4)*5)) = -10
//         (((2*3)-4)*5) = 10
// 思想：
// 根据操作将表达式分割为左右两个子表达式，然后分别递归求解子表达式。
func diffWaysToCompute(input string) []int {
	temp := make(map[string][]int, 0) // 存放每个表达式的结果
	return ways(input, temp)
}

func ways(input string, temp map[string][]int) []int {
	if _, ok := temp[input]; ok {
		return temp[input]
	}

	ans := make([]int, 0)
	for i := 0; i < len(input); i++ { // 循环表达式
		op := input[i]
		if op == '+' || op == '-' || op == '*' { // 如果是操作符，则分割表达式
			left := input[:i]
			right := input[i+1:]
			leftAns := ways(left, temp) // 递归
			rightAns := ways(right, temp)

			for _, l := range leftAns { // 将左右子表达式的解合并
				for _, r := range rightAns {
					switch op {
					case '+':
						ans = append(ans, l+r)
					case '-':
						ans = append(ans, l-r)
					case '*':
						ans = append(ans, l*r)
					}
				}
			}

		}
	}

	if len(ans) == 0 {
		i, _ := strconv.Atoi(input)
		ans = append(ans, i)
	}
	temp[input] = ans

	return temp[input]
}

// 下面是该题leetcode排名最靠前的牛逼代码
func isOp(c byte) bool {
	return c == '+' || c == '-' || c == '*'
}

func splitOp(input string) ([]int, []byte) {
	nums := make([]int, 0)
	ops := make([]byte, 0)

	for i := 0; i < len(input); {
		if isOp(input[i]) {
			ops = append(ops, input[i])
			i++
		} else {
			j := i + 1
			for ; j < len(input) && !isOp(input[j]); j++ {
			}
			temp, _ := strconv.ParseInt(input[i:j], 10, 64)
			nums = append(nums, int(temp))
			i = j
		}
	}

	return nums, ops
}

func opCompute(first int, op byte, second int) int {
	switch op {
	case '+':
		return first + second
	case '-':
		return first - second
	case '*':
		return first * second
	case '/':
		return first / second
	}
	return 0
}

func compute(nums []int, ops []byte) []int {
	if len(nums) == 0 {
		return []int{}
	}
	if len(nums) == 1 {
		return []int{nums[0]}
	}
	if len(nums) == 2 {
		return []int{opCompute(nums[0], ops[0], nums[1])}
	}
	res := make([]int, 0)
	for i := range ops {
		formerNums := nums[:i+1]
		formerOps := ops[:i]

		latterNums := nums[i+1:]
		latterOps := ops[i+1:]

		formerRes := compute(formerNums, formerOps)
		latterRes := compute(latterNums, latterOps)
		for _, former := range formerRes {
			for _, latter := range latterRes {
				res = append(res, opCompute(former, ops[i], latter))
			}
		}
	}
	return res
}

func diffWaysToCompute2(input string) []int {
	nums, ops := splitOp(input)
	return compute(nums, ops)
}
