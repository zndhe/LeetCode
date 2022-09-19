/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 */

// @lc code=start
func search(nums []int, target int) bool {
	n := len(nums)
	left, right := 0, n-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return true
		}
		if nums[left] == nums[mid] {
			left++
		} else if nums[left] > nums[mid] {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid
			}
		} else {
			if target < nums[mid] && target >= nums[left] {
				right = mid
			} else {
				left = mid + 1
			}
		}
	}
	if nums[left] == target {
		return true
	}
	return false
}

// @lc code=end

