/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[left] > nums[mid] {
			if target > nums[mid] && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if nums[mid] > target && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}

// @lc code=end

