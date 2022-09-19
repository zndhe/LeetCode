/*
 * @lc app=leetcode.cn id=547 lang=golang
 *
 * [547] 省份数量
 */

// @lc code=start
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	visited := make([]bool, n)
	res := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(isConnected, i, visited)
			res++
		}
	}
	return res
}

func dfs(isConnected [][]int, i int, visited []bool) {
	visited[i] = true
	for k := 0; k < len(isConnected); k++ {
		if isConnected[i][k] == 1 && !visited[k] {
			dfs(isConnected, k, visited)
		}
	}
}

// @lc code=end

