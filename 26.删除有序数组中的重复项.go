/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	n := len(nums)
	left := 1 //
	for i := 1; i < n; i++ {
		if nums[i] != nums[i-1] {
			nums[left] = nums[i] //
			left++               //
		}
	}
	return left
}

// @lc code=end

