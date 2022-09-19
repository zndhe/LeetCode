/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
	var ans [][]int
	var used []int
	var backtracking func(int, int)

	backtracking = func(level, count int) {
		if count == 0 {
			res := make([]int, k)
			copy(res, used)
			ans = append(ans, res)
			return
		}
		for i := level; count > 0 && i <= n; i++ {
			used = append(used, i)
			//不考虑顺序，故不用考虑 1 3 回去找2，直接往下遍历即可
			//如果考虑，则每次都需要从0遍历，通过一个used数组判断前面序列已用过的数字
			backtracking(i+1, count-1)
			used = used[:len(used)-1]
		}
	}

	backtracking(1, k)
	return ans
}

// @lc code=end

