package greedy

import "fmt"

var (
	MONEY = [6]int{200, 100, 50, 10, 5, 1}
	NUM   = 6
)

// 1. 有1元、5元、10元、50元、100元、200元的钞票无穷多张，现在用这些钞票支付X元，最少需要多少张钞票？
func Pay(x int) int {
	var count int
	for i := 0; i < NUM && x > 0; i++ {
		use := x / MONEY[i]  // 需要面值为MONEY[i]的use张
		count = count + use  // 总共增加use张
		x = x - use*MONEY[i] // 从总金额中减去已使用的金额
		fmt.Printf("需要面额%d的%d张\n", MONEY[i], use)
	}
	return count
}
