/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	res := math.MaxInt16
	//上个基准数字，避免重复组合的出现
	a := 0
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == a {
			continue
		}
		a = nums[i]
		left, right := i+1, n-1
		for left < right {
			b := nums[left]
			c := nums[right]
			sum := a + b + c
			if sum == target {
				return sum
			} else if sum < target {
				left++
			} else {
				right--
			}
			if abs(target, sum) < abs(target, res) {
				res = sum
			}
		}
	}
	return res
}

func abs(a int, b int) int {
	if a >= b {
		return a - b
	} else {
		return b - a
	}
}

// @lc code=end

