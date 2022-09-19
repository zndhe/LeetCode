/*
 * @lc app=leetcode.cn id=406 lang=golang
 *
 * [406] 根据身高重建队列
 */

// @lc code=start
//从低到高
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] < people[j][0] || people[i][0] == people[j][0] && people[i][1] > people[j][1]
	})
	n := len(people)
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		space := people[i][1] + 1
		for j := 0; j < n; j++ {
			if res[j] == nil {
				space--
				if space == 0 {
					res[j] = people[i]
					break
				}
			}
		}
	}
	return res
}
// @lc code=end

