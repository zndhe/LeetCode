/*
 * @lc app=leetcode.cn id=504 lang=golang
 *
 * [504] 七进制数
 */

// @lc code=start
func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	neg := false
	if num < 0 {
		neg = true
		num = -num
	}
	res := ""
	nums := []int{}
	for num != 0 {
		rem := num % 7
		nums = append(nums, rem)
		num /= 7
	}
	if neg {
		res += "-"
	}
	for i := len(nums) - 1; i >= 0; i-- {
		res += strconv.Itoa(nums[i])
	}

	return res
}

// @lc code=end

