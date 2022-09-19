/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func romanToInt(s string) int {
	n := len(s)
	res := 0
	for i := range s{
		num := getValue(s[i])
		if i < n-1 && num < getValue(s[i+1]){
			res -= num
		} else {
			res += num
		}
	}
	return res
}
func getValue(b byte) (res int) {
	switch b {
	case 'I':
		res = 1
	case 'V':
		res = 5
	case 'X':
		res = 10
	case 'L':
		res = 50
	case 'C':
		res = 100
	case 'D':
		res = 500
	case 'M':
		res = 1000
	}
	return
}
// @lc code=end

