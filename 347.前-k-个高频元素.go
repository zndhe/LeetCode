/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	numsMap := map[int]int{}
	maxVal := 0
	for _, v := range nums {
		numsMap[v]++
		maxVal = max(maxVal, numsMap[v])
	}
	values := make([][]int, maxVal+1)
	for k, v := range numsMap {
		values[v] = append(values[v], k)
	}
	res := []int{}
	for i := len(values) - 1; i > 0 && len(res) < k; i-- {
		for _, v := range values[i] {
			res = append(res, v)
		}
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

// @lc code=end

