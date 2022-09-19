/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */

// @lc code=start

func numDecodings(s string) int {
	n := len(s)
	a, b, c := 0, 1, 0
	for i := 1; i <= n; i++ {
		c = 0
		if s[i-1] != '0' {
			c += b
		}
		if i > 1 && s[i-2] != '0' && (s[i-2]-'0')*10+s[i-1]-'0' <= 26 {
			c += a
		}
		a, b = b, c
	}
	return c
}

// func numDecodings(s string) int {
// 	n := len(s)
// 	dp := make([]int, n+1)
// 	dp[0] = 1
// 	for i := 1; i <= n; i++ {
// 		if s[i-1] != '0' {
// 			dp[i] += dp[i-1]
// 		}
// 		if i > 1 && s[i-2] != '0' && (s[i-2]-'0')*10+s[i-1]-'0' <= 26 {
// 			dp[i] += dp[i-2]
// 		}
// 	}
// 	return dp[n]
// }

// @lc code=end

