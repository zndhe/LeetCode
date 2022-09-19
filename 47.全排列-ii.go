/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	path := []int{}
	vis := make([]bool, n)
	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == n {
			ans = append(ans, append([]int(nil), path...))
			return

		}
		for i, v := range nums {
			//因为是从头遍历，当vis[i-1]==false
			//说明前面的已经遍历一次了
			//例如[1,1,2] 在第一轮遍历结束，读取第二个1时，得到的数组与第一轮一致
			//此时第一轮结束后vis[0]已经重置为false
			if vis[i] || i > 0 && !vis[i-1] && v == nums[i-1] {
				continue
			}
			path = append(path, v)
			vis[i] = true
			backtrack(idx + 1)
			vis[i] = false
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return
}

// @lc code=end

