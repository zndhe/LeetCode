/*
 * @lc app=leetcode.cn id=665 lang=golang
 *
 * [665] 非递减数列
 */

// @lc code=start
func checkPossibility(nums []int) bool {
	n := len(nums)
	sub := 1
	for i := 0; i < n-1; i++ {
		if nums[i+1] < nums[i] {
			sub--
			if i > 0 && nums[i+1] < nums[i-1] {
				nums[i+1] = nums[i]
			}
			if sub < 0 {
				return false
			}
		}
	}
	return true
}

// @lc code=end

