/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
// var next = [4][2]int{
// 	{1, 0},
// 	{-1, 0},
// 	{0, 1},
// 	{0, -1},
// }

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	var backtracking func(int, int, int) bool
	backtracking = func(i, j, k int) bool {
		if i < 0 || j < 0 || i >= m || j >= n || word[k] != board[i][j] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		board[i][j] -= 60
		// if backtracking(i+v[0], j+v[1], k+1) {
		// 	return true
		// }
		if backtracking(i+1, j, k+1) {
			return true
		}
		if backtracking(i-1, j, k+1) {
			return true
		}
		if backtracking(i, j+1, k+1) {
			return true
		}
		if backtracking(i, j-1, k+1) {
			return true
		}
		//还原
		board[i][j] += 60
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backtracking(i, j, 0) {
				return true
			}
		}
	}
	return false
}

// @lc code=end

