/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 */

// @lc code=start
func maxProfit(k int, prices []int) int {
	day := len(prices)
	if day < 2 || k < 1 {
		return 0
	}
	//buy sell分别是在买卖后剩下的钱
	sell, buy := make([]int, k+1), make([]int, k+1)
	for i := 0; i <= k; i++ {
		buy[i] = math.MinInt32
	}
	for i := 0; i < day; i++ {
		for j := 1; j <= k; j++ {
			buy[j] = max(buy[j], sell[j-1]-prices[i])
			sell[j] = max(sell[j], buy[j]+prices[i])
		}
	}
	return sell[k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

