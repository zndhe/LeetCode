/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	pre1, pre2 := 1, 2
	cur := 0
	//状态转移方程dp[i] = dp[i-1] + dp[i-2]。
	for i := 2; i < n; i++ {
		cur = pre1 + pre2
		pre1 = pre2
		pre2 = cur
	}
	return cur
}

// @lc code=end

