/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	l, r, target := 0, len(nums)-1, len(nums)-k
	for l < r {
		mid := quickSort(nums, l, r)
		if mid == target {
			return nums[mid]
		} else if mid > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return nums[l]
}
func quickSort(nums []int, l, r int) int {
	left, right := l, r
	base := nums[left]
	for left < right {
		//去等号会在==base时停下，但是需要考虑到left right 指向同个数
		//在互换时将left+1 right-1
		for left < right && nums[right] >= base {
			right--
		}
		for left < right && nums[left] <= base {
			left++
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
		}

	}
	nums[l], nums[left] = nums[left], nums[l]
	return left
}

// @lc code=end

