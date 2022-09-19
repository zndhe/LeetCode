/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	i, sum, sign, n := 0, 0, 1, len(s)
	for i < n && s[i] == ' ' {
		i++
	}
	if i < n {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			sign = 1
			i++
		}
	}
	for i < n && s[i] >= '0' && s[i] <= '9' {
		sum = sum*10 + int(s[i]-'0')
		if sum * sign > math.MaxInt32 {
			return math.MaxInt32
		}else if sum * sign < math.MinInt32 {
			return math.MinInt32
		}
		i++
	}
	return sum * sign
}

// @lc code=end

