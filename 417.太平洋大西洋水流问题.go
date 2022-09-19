/*
 * @lc app=leetcode.cn id=417 lang=golang
 *
 * [417] 太平洋大西洋水流问题
 */

// @lc code=start
var next = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func pacificAtlantic(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])
	res := [][]int{}
	pacific := make([][]bool, m)
	atlantic := make([][]bool, m)
	for i := range pacific {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}

	var dfs func([][]bool, int, int)
	dfs = func(ocean [][]bool, x, y int) {
		if ocean[x][y] {
			return
		}
		ocean[x][y] = true
		for _, v := range next {
			if nx, ny := x+v.x, y+v.y; 0 <= nx && nx < m && 0 <= ny && ny < n && heights[x][y] <= heights[nx][ny] {
				dfs(ocean, nx, ny)
			}
		}
	}
	//左右
	for i := 0; i < m; i++ {
		dfs(pacific, i, 0)
		dfs(atlantic, i, n-1)
	}
	//上下
	for j := 0; j < n; j++ {
		dfs(pacific, 0, j)
		dfs(atlantic, m-1, j)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pacific[i][j] && atlantic[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}

	return res
}

// @lc code=end

