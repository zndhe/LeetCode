/*
 * @lc app=leetcode.cn id=1031 lang=golang
 *
 * [1031] 两个非重叠子数组的最大和
 */

// @lc code=start
func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	n := len(nums)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + nums[i-1]
	}
	res := 0
	for j := firstLen; j <= n; j++ {
		sum1 := dp[j] - dp[j-firstLen]
		sum2 := 0
		for k := secondLen; k <= j-firstLen; k++ {
			sum2 = max(sum2, dp[k]-dp[k-secondLen])
		}
		for k := j + secondLen; k <= n; k++ {
			sum2 = max(sum2, dp[k]-dp[k-secondLen])
		}
		res = max(res, sum1+sum2)
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

