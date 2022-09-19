/*
 * @lc app=leetcode.cn id=435 lang=golang
 *
 * [435] 无重叠区间
 */

// @lc code=start
func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	res, right := 0, intervals[0][1]
	for i := 1; i < n; i++ {
		if intervals[i][0] >= right {
			right = intervals[i][1]
		} else {
			res++
		}
	}
	return res
}

// @lc code=end

