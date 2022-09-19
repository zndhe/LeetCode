/*
 * @lc app=leetcode.cn id=524 lang=golang
 *
 * [524] 通过删除字母匹配到字典里最长单词
 */

// @lc code=start
func findLongestWord(s string, dictionary []string) string {
	sort.Slice(dictionary, func(i, j int) bool {
		a, b := dictionary[i], dictionary[j]
		return len(a) > len(b) || len(a) == len(b) && a < b
	})
	for _, t := range dictionary {
		i := 0
		for j, _ := range s {
			if len(s) - j < len(t) - i {
				break
			}
			if s[j] == t[i] {
				i++
				if i == len(t) {
					return t
				}
			}
		}
	}
	return ""
}

// @lc code=end

