/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 */

// @lc code=start
func maxProfit(prices []int) int {
	day := len(prices)
	if day < 2 {
		return 0
	}
	//buy sell分别是在买卖后剩下的钱
	sell, buy := make([]int, day), make([]int, day)
	for i := 0; i < day; i++ {
		buy[i] = math.MinInt32
	}
	buy[0] = -prices[0]
	buy[1] = max(buy[0], -prices[1])
	sell[1] = max(sell[0], buy[1]+prices[1])
	for i := 2; i < day; i++ {
		buy[i] = max(buy[i-1], sell[i-2]-prices[i])
		sell[i] = max(sell[i-1], buy[i]+prices[i])
	}
	return sell[day-1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

