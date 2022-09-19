/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	ans := []string{}
	backtracking(n, n, "", &ans)
	return ans
}

func backtracking(lindex, rindex int, str string, ans *[]string) {
	if lindex == 0 && rindex == 0 {
		*ans = append(*ans, str)
		return
	}
	if lindex > 0 {
		backtracking(lindex-1, rindex, str+"(", ans)
	}
	if rindex > 0 && lindex < rindex{
		backtracking(lindex, rindex-1, str+")", ans)
	}
}

// @lc code=end

