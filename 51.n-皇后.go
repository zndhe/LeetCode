/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 */

// @lc code=start
var ans [][]string

func solveNQueens(n int) [][]string {
	ans = [][]string{}
	queens := make([]int, n)
	column, rdiag, ldiag := make([]bool, n), make([]bool, 2*n-1), make([]bool, 2*n-1)
	backtracking(0, n, column, rdiag, ldiag, queens)
	return ans
}

func backtracking(row, n int, column, rdiag, ldiag []bool, queens []int) {
	if row == n {
		board := generateBoard(queens, n)
		ans = append(ans, board)
		return
	}
	for i := 0; i < n; i++ {
		if column[i] || rdiag[row+i] || ldiag[n-row+i-1] {
			continue
		}
		column[i], rdiag[row+i], ldiag[n-row+i-1] = true, true, true
		queens[row] = i
		backtracking(row+1, n, column, rdiag, ldiag, queens)
		queens[row] = -1
		column[i], rdiag[row+i], ldiag[n-row+i-1] = false, false, false
	}
}

func generateBoard(queens []int, n int) []string {
	board := []string{}
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		//整行赋值—，然后通过index改变为Q
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
}

// @lc code=end

