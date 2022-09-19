/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	if s == " " {
		return 1
	}
	n := len(s)
	j := 0
	cmap := map[byte]int{}
	res := 0
	//从每个i开始算长度，取最长
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(cmap, s[i-1])
		}
		for j < n && cmap[s[j]] == 0 {
			cmap[s[j]]++
			j++
		}
		res = max(res, j-i)
	}
	return res
}
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end

