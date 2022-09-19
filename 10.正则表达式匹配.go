/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

// @lc code=start
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	f := make([][]bool, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n+1)
	}
	f[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				//先与两位前的bool比较，因为*可以取0个
				f[i][j] = f[i][j-2]
				if matches(i, j-1) {
					//如果 s[i-1] 与 p[j-2] 匹配，
					//因为有多个重复的字母，此时是看列保证前面的序列都是匹配
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else if matches(i, j) {
				//保证前面的序列都匹配，取斜对角的取值
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
		}
	}
	return f[m][n]
}

// @lc code=end

