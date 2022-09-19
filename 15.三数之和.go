/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	res := make([][]int, 0)
	//上个基准数字，避免重复组合的出现
	a := 0
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return res
		}
		if (i > 0 && nums[i] == a) || nums[i]+nums[n-1]+nums[n-2] < 0 {
			continue
		}
		a = nums[i]
		//避免重复
		left, right := i+1, n-1
		for left < right {
			b := nums[left]
			c := nums[right]
			sum := a + b + c
			if sum == 0 {
				res = append(res, []int{a, b, c})
				for {
					left++
					if left >= n-1 || nums[left] != b {
						break
					}
				}
				for {
					right--
					if right <= i || nums[right] != c {
						break
					}
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}

// @lc code=end

