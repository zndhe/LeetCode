/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 */

// @lc code=start
func solveSudoku(board [][]byte) {
	var rows, columns, block [9][9]bool
	spaces := [][2]int{}
	//赋值
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				spaces = append(spaces, [2]int{i, j})
			} else {
				num := board[i][j] - '1'
				rows[i][num] = true
				columns[j][num] = true
				block[i/3*3+j/3][num] = true
			}
		}
	}
	var backtracking func(int) bool
	backtracking = func(pos int) bool {
		if pos == len(spaces) {
			return true
		}
		i, j := spaces[pos][0], spaces[pos][1]
		for k := byte(0); k < 9; k++ {
			if !rows[i][k] && !columns[j][k] && !block[i/3*3+j/3][k] {
				rows[i][k] = true
				columns[j][k] = true
				block[i/3*3+j/3][k] = true
				board[i][j] = k + '1'
				if backtracking(pos + 1) {
					return true
				}
				rows[i][k] = false
				columns[j][k] = false
				block[i/3*3+j/3][k] = false
			}
		}
		return false
	}
	backtracking(0)

}

// @lc code=end

