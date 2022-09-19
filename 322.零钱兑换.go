/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	sort.Ints(coins)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		// for _, v := range coins {
		// 	if i >= v {
		// 		dp[i] = min(dp[i], dp[i-v]+1)
		// 	}
		// }
		for j := 0; j < len(coins) && i-coins[j] >= 0; j++ {
			dp[i] = min(dp[i], dp[i-coins[j]]+1)
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

