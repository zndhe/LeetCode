/*
 * @lc app=leetcode.cn id=650 lang=golang
 *
 * [650] 只有两个键的键盘
 */

// @lc code=start
func minSteps(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = i
		for j := 1; j*j <= i; j++ {
			if i%j == 0 {
				if i/j > j {
					dp[i] = min(dp[i], dp[i/j]+j)
				} else {
					dp[i] = min(dp[i], dp[j]+i/j)
				}
			}
		}
	}
	return dp[n]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

