/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	diff := sum - target
	if diff < 0 || diff%2 == 1 {
		return 0
	}
	//减的总值
	diff /= 2
	//前i个数组成j的组合数(i通过优化空间)
	dp := make([]int, diff+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		for j := diff; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}
	return dp[diff]
}

// @lc code=end

