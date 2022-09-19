/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for _,v := range wordDict {
			len := len(v)
			if i >= len && s[i-len:i] == v {
				// i-len - len 确定是有对应的，找到 i-len之前的
				//如果是对的，则0-len都是对的
				dp[i] =  dp[i] || dp[i-len]
			}
		}
	}
	return dp[n]
}

// @lc code=end

