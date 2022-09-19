/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
func convert(s string, numRows int) string {
	if numRows == 1{
		return s
	}
	ans := make([][]byte, numRows)
	r := 0
	time := 2*numRows - 2
	for k, v := range s {
		ans[r] = append(ans[r], byte(v))
		if k%time < numRows-1 {
			r++
		} else {
			r--
		}
	}
	return string(bytes.Join(ans,nil))
}

// @lc code=end

