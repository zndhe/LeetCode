/*
 * @lc app=leetcode.cn id=542 lang=golang
 *
 * [542] 01 矩阵
 */

// @lc code=start
func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				continue
			} else {
				//这时候的0可能是初始化出来的，并不是真的距离0
				left, top := math.MaxInt16, math.MaxInt16
				if i > 0 {
					top = dp[i-1][j] + 1
				}
				if j > 0 {
					left = dp[i][j-1] + 1
				}
				dp[i][j] = min(left, top)
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if mat[i][j] != 0 {
				if i < m-1 {
					dp[i][j] = min(dp[i][j], dp[i+1][j]+1)
				}
				if j < n-1 {
					dp[i][j] = min(dp[i][j], dp[i][j+1]+1)
				}
			}
		}
	}
	return dp
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

