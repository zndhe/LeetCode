/*
 * @lc app=leetcode.cn id=763 lang=golang
 *
 * [763] 划分字母区间
 */

// @lc code=start
func partitionLabels(s string) []int {
	lastPos := [26]int{}
	var res []int
	n := len(s)
	for i := 0; i < n; i++ {
		lastPos[s[i]-'a'] = i
	}
	left, right := 0, 0
	for i := 0; i < n; i++ {
		if lastPos[s[i]-'a'] > right {
			right = lastPos[s[i]-'a']
		}
		if i == right {
			res = append(res, right-left+1)
			left = right + 1
		}
	}
	return res
}

// @lc code=end

