/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	res := make([][]int, 0)
	for j := 0; j < n-3 && nums[j]+nums[j+1]+nums[j+2]+nums[j+3] <= target; j++ {
		if (j > 0 && nums[j] == nums[j-1]) || nums[j]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for i := j + 1; i < n-2 && nums[j]+nums[i]+nums[i+1]+nums[i+2] <= target; i++ {
			if (i > j+1 && nums[i] == nums[i-1]) || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			left, right := i+1, n-1
			for left < right {
				sum := nums[j] + nums[i] + nums[left] + nums[right]
				if sum == target {
					res = append(res, []int{nums[j], nums[i], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {

					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {

					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return res
}

// @lc code=end

