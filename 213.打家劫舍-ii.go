/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	pre1, pre2 := 0, 0
	cur1 := 0
	//状态转移方程为 dp[i] = max(dp[i-1],nums[i] + dp[i-2])
	for i := 0; i < n-1; i++ {
		cur1 = max(pre1+nums[i], pre2)
		pre1 = pre2
		pre2 = cur1
	}
	pre1, pre2 = 0, 0
	cur2 := 0
	//状态转移方程为 dp[i] = max(dp[i-1],nums[i] + dp[i-2])
	for i := 1; i < n; i++ {
		cur2 = max(pre1+nums[i], pre2)
		pre1 = pre2
		pre2 = cur2
	}
	return max(cur1, cur2)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
// @lc code=end

