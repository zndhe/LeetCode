/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	pre1, pre2 := 0, 0
	cur := 0
	//状态转移方程为 dp[i] = max(dp[i-1],nums[i] + dp[i-2])
	for i := 0; i < n; i++ {
		cur = max(pre1+nums[i], pre2)
		pre1 = pre2
		pre2 = cur
	}
	return cur
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end

