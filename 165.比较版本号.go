/*
 * @lc app=leetcode.cn id=165 lang=golang
 *
 * [165] 比较版本号
 */

// @lc code=start
func compareVersion(version1 string, version2 string) int {
	n, m := len(version1), len(version2)
	i, j := 0, 0
	for i < n || j < m {
		a := 0
		for ; i < n && version1[i] != '.'; i++ {
			a = a*10 + int(version1[i]-'0')
		}
		i++
		b := 0
		for ; j < m && version2[j] != '.'; j++ {
			b = b*10 + int(version2[j]-'0')
		}
		j++
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
	}
	return 0
}

// @lc code=end

