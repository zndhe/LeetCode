/*
 * @lc app=leetcode.cn id=452 lang=golang
 *
 * [452] 用最少数量的箭引爆气球
 */

// @lc code=start
// func findMinArrowShots(points [][]int) int {
// 	sort.Slice(points, func(i, j int) bool {
// 		return points[i][1] < points[j][1]
// 	})
// 	n := len(points)
// 	res, right := 1, points[0][1]
// 	for i := 1; i < n; i++ {
// 		if points[i][0] > right {
// 			right = points[i][1]
// 			res++
// 		}
// 	}
// 	return res
// }

//sort.Slice用到反射的函数，稍慢
type SortBy [][]int

func (a SortBy) Len() int {
	return len(a)
}
func (a SortBy) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a SortBy) Less(i, j int) bool {
	return a[i][1] < a[j][1]
}
func findMinArrowShots(points [][]int) int {
	res := 1
	sort.Sort(SortBy(points))
	lastEnd := points[0][1]
	for _, point := range points {
		if point[0] > lastEnd {
			res++
			lastEnd = point[1]
		}
	}
	return res
}


// @lc code=end

