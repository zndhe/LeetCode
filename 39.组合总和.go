/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	temp, res := []int{}, [][]int{}
	sort.Ints(candidates)
	backtracking(candidates, target, 0, temp, &res)
	return res
}
func backtracking(candidates []int, target int, index int, temp []int, res *[][]int) {
	if target <= 0 {
		if target == 0 {
			a := make([]int, len(temp))
			copy(a, temp)
			*res = append(*res, a)
		}
		return
	}
	for i := index; i < len(candidates); i++ {
		if candidates[i] > target {
			break
		}
		temp = append(temp, candidates[i])
		backtracking(candidates, target-candidates[i], i, temp, res)
		temp = temp[:len(temp)-1]
	}
}

// @lc code=end

