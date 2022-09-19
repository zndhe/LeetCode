/*
 * @lc app=leetcode.cn id=934 lang=golang
 *
 * [934] 最短的桥
 */

// @lc code=start
var next = [4][2]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func shortestBridge(grid [][]int) int {
	queue := make([][]int, 0)
	first := true
	for i := 0; i < len(grid) && first; i++ {
		for j := 0; j < len(grid[0]) && first; j++ {
			if grid[i][j] == 1 {
				//深搜寻找第一个岛
				dfs(grid, &queue, i, j)
				first = false
			}
		}
	}
	//在第一个岛的基础上广搜直到接触到第二个岛
	return bfs(grid, queue)
}

func dfs(grid [][]int, queue *[][]int, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}
	if grid[i][j] == 0 {
		*queue = append(*queue, []int{i, j})
	} else if grid[i][j] == 1 {
		grid[i][j] = 2
		dfs(grid, queue, i-1, j)
		dfs(grid, queue, i+1, j)
		dfs(grid, queue, i, j-1)
		dfs(grid, queue, i, j+1)
	}
}

func bfs(grid [][]int, queue [][]int) int {
	ans := 0
	for len(queue) > 0 {
		ans++
		for k := len(queue); k > 0; k-- {
			i, j := queue[0][0], queue[0][1]
			queue = queue[1:]
			for _, v := range next {
				newI, newJ := i+v[0], j+v[1]
				if newI >= 0 && newI < len(grid) && newJ >= 0 && newJ < len(grid[0]) {
					if grid[newI][newJ] == 1 {
						return ans
					}
					if grid[newI][newJ] == 2 {
						continue
					}
					grid[newI][newJ] = 2
					queue = append(queue, []int{newI, newJ})
				}
			}
		}
	}
	return ans
}

// @lc code=end

