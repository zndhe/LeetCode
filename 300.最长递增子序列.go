/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	dp := []int{}
	for _, v := range nums {
		i := sort.SearchInts(dp, v)
		if i == len(dp) {
			dp = append(dp, v)
		} else {
			dp[i] = v		
		}
	}
	return len(dp)
}

// @lc code=end

