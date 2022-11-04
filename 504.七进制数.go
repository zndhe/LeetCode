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
	for num != 0 {
		res = strconv.Itoa(num%7) + res
		num /= 7
	}
	if neg {
		res = "-" + res
	}

	return res
}

// @lc code=end

