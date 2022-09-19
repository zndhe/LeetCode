/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
func minWindow(s string, t string) string {
	ori := make([]int, 128)
	cnt := make([]int, 128)
	for _, v := range t {
		ori[v]++
	}

	distance := 0
	sLen := len(s)
	minLen := math.MaxInt32
	ansL := -1

	for l, r := 0, 0; r < sLen; r++ {
		if ori[s[r]] > 0 {
			if cnt[s[r]] < ori[s[r]] {
				distance++
			}
			cnt[s[r]]++
		}
		//现在的lr之间序列满足t，尝试移动l
		for distance == len(t) {
			if r-l+1 < minLen {
				minLen = r - l + 1
				ansL = l
			}
			if ori[s[l]] > 0 {
				if cnt[s[l]] == ori[s[l]] {
					distance--
				}
				cnt[s[l]]--
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL : ansL+minLen]
}

// @lc code=end

