/*
 * @lc app=leetcode.cn id=646 lang=golang
 *
 * [646] 最长数对链
 */

// @lc code=start

//动归
// func findLongestChain(pairs [][]int) int {
// 	dp := []int{}
// 	sort.Slice(pairs, func(i, j int) bool {
// 		return pairs[i][0] < pairs[j][0]
// 	})
// 	for _, v := range pairs {
// 		// fmt.Println(dp, v[0])
// 		i := sort.SearchInts(dp, v[0])
// 		// fmt.Println(i)
// 		if i == len(dp) {
// 			dp = append(dp, v[1])
// 		} else if v[1] < dp[i] {
// 			dp[i] = v[1]
// 		}
// 	}
// 	return len(dp)
// }

//贪心 ：无重叠区间，边界也不重叠
func findLongestChain(pairs [][]int) int {
	res := 0
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})
	// first := pairs[0][0]
	end := math.MinInt32
	for i := 0; i < len(pairs); i++ {
		if pairs[i][0] > end {
			res++
			end = pairs[i][1]
		}
	}
	return res
}

// @lc code=end

