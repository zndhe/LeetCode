/*
 * @lc app=leetcode.cn id=135 lang=golang
 *
 * [135] 分发糖果
 */

// @lc code=start
func candy(ratings []int) int {
	n := len(ratings)
	if n < 2 {
		return n
	}

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		if i > 0 && ratings[i] > ratings[i-1] {
			nums[i] = nums[i-1] + 1
		} else {
			nums[i] = 1
		}
	}
	res := nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			nums[i] = max(nums[i], nums[i+1]+1)
		}
		res += nums[i]
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

