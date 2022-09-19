3/*
 * @lc app=leetcode.cn id=413 lang=golang
 *
 * [413] 等差数列划分
 */

// @lc code=start
func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}
	ans := 0
	count := 0
	for i := 2; i < n; i++ {
		if nums[i] - nums[i-1] == nums[i-1] - nums[i-2] {
			count++
		} else {
			count = 0
		}
		ans += count
	}
	return ans
}
// @lc code=end

