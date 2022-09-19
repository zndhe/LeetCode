/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
	l, r, sum := 0, len(numbers)-1, 0
	for l < r {
		sum = numbers[l] + numbers[r]
		if sum == target {
			return []int{l+1, r+1}
		} else if sum > target {
			r--
		} else {
			l++
		}
	}
	return []int{-1, -1}
}

// @lc code=end

