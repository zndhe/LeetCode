/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	ans := [][]int{}
	used := make([]bool, len(candidates))
	path := []int{}
	backtracking(&ans, used, path, candidates, 0, target)
	return ans
}

func backtracking(ans *[][]int, used []bool, path, candidates []int, level, target int) {
	if target <= 0 {
		if target == 0 {
			path = append(path[:0:0], path...)
			*ans = append(*ans, path)
		}
		return
	}
	for i := level; i < len(candidates); i++ {
		//处理重复的核心代码
		if i > 0 && !used[i-1] && candidates[i] == candidates[i-1] {
			continue
		}
		path = append(path, candidates[i])
		used[i] = true
		backtracking(ans, used, path, candidates, i+1, target-candidates[i])
		used[i] = false
		path = path[:len(path)-1]
	}
}

// @lc code=end

