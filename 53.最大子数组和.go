/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	pre, res, cur := nums[0], nums[0], 0
	for i := 1; i < len(nums); i++ {
		if pre > 0 {
			cur = nums[i] + pre
		} else {
			cur = nums[i]
		}
		pre = cur
		res = max(res, cur)
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

