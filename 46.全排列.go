/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	ans := [][]int{}
	var backtracking func([]int, int)
	backtracking = func(nums []int, level int) {
		if level == len(nums)-1 {
			res := make([]int, len(nums))
			copy(res, nums)
			ans = append(ans, res)
			return
		}
		for i := level; i < len(nums); i++ {
			nums[i], nums[level] = nums[level], nums[i]
			backtracking(nums, level+1)
			nums[i], nums[level] = nums[level], nums[i]
		}
	}
	backtracking(nums, 0)
	return ans
}

// @lc code=end

