/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for i := 0; i < len(dp[0]); i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] != word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i-1][j-1]
			}
			//为什么考虑两侧，例:a,b,c,d  g,a,b,c,dz
			dp[i][j] = minss(dp[i][j], dp[i-1][j]+1, dp[i][j-1]+1)
		}
	}
	return dp[m][n]
}

func minss(a int, b ...int) int {
	for _, i := range b {
		if i < a {
			a = i
		}
	}
	return a
}

// @lc code=end

