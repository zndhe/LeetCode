/*
 * @lc app=leetcode.cn id=528 lang=golang
 *
 * [528] 按权重随机选择
 */

// @lc code=start
type Solution struct {
	prefixSum []int
}

func Constructor(w []int) Solution {
	prefixSum := make([]int, len(w))
	for i, v := range w {
		if i == 0 {
			prefixSum[0] = v
			continue
		}
		prefixSum[i] = prefixSum[i-1] + v
	}
	return Solution{
		prefixSum: prefixSum,
	}
}

func (this *Solution) PickIndex() int {
	n := rand.Intn(this.prefixSum[len(this.prefixSum)-1]) + 1
	low, high := 0, len(this.prefixSum)-1
	for low < high {
		mid := (low + high) / 2
		if this.prefixSum[mid] == n {
			return mid
		} else if this.prefixSum[mid] < n {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
// @lc code=end

