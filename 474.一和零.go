/*
 * @lc app=leetcode.cn id=474 lang=golang
 *
 * [474] 一和零
 */

// @lc code=start
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for _, s := range strs {
		zero := strings.Count(s,"0")
		one := len(s) - zero
		// zero, one := count(s)
		for i := m; i >= zero; i--{
			for j := n; j >= one; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zero][j-one]+1)
			}
		}
	}
	return dp[m][n]
}
func count(s string) (int, int) {
	zero, one := len(s), 0
	for _, v := range s {
		if v == '1' {
			zero--
			one++
		}
	}
	return zero, one
}

func max(a , b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

