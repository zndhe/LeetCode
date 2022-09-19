/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	leftIdx := binarySearch(nums, target)
	rightIdx := binarySearch(nums, target+1) - 1
	if leftIdx <= rightIdx {
		return []int{leftIdx, rightIdx}
	}
	return []int{-1, -1}
}
func binarySearch(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	res := n
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid - 1
			res = mid
		} else {
			left = mid + 1
		}
	}
	return res
}

// @lc code=end

