/*
 * @lc app=leetcode.cn id=932 lang=golang
 *
 * [932] 漂亮数组
 */

// @lc code=start
func beautifulArray(n int) []int {
	dp := make([][]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = make([]int, 0, i)

		if i == 1 {
			dp[i] = append(dp[i], 1)
			continue
		}
		odds, evens := (i+1)>>1, i>>1
		for _, v := range dp[odds] {
			dp[i] = append(dp[i], v*2-1)
		}
		for _, v := range dp[evens] {
			dp[i] = append(dp[i], v*2)
		}
	}
	return dp[n]
}

func beautifulArray1(n int) []int {
	//分治，不用记忆式搜索
	//记忆式是空间换时间，这道题直接0ms
	if n == 1 {
		return []int{1}
	}
	array := make([]int, 0, n)
	odds := beautifulArray((n + 1) >> 1)
	evens := beautifulArray(n >> 1)

	for _, v := range odds {
		array = append(array, 2*v-1)
	}
	for _, v := range evens {
		array = append(array, 2*v)
	}

	return array
}

// @lc code=end

