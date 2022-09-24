/*
 * @lc app=leetcode.cn id=312 lang=golang
 *
 * [312] 戳气球
 */

// @lc code=start
func maxCoins(nums []int) int {
	n := len(nums)
	val := make([]int, n+2)
	val[0], val[n+1] = 1, 1
	for i := 1; i <= n; i++ {
		val[i] = nums[i-1]
	}
	res := make([][]int, n+2)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, n+2)
		for j := 0; j < len(res[i]); j++ {
			res[i][j] = -1
		}
	}
	return solve(0, n+1, val, res)
}

func solve(left, right int, val []int, res [][]int) int {
	if left >= right-1 {
		return 0
	}
	if res[left][right] != -1 {
		return res[left][right]
	}
	for i := left + 1; i < right; i++ {
		sum := val[left] * val[i] * val[right]
		sum += solve(left, i, val, res) + solve(i, right, val, res)
		res[left][right] = max(res[left][right], sum)
	}
	return res[left][right]
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

