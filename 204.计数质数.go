/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 */

// @lc code=start
func countPrimes(n int) int {
	count := 0
	isNotPrime := make([]bool, n)
	for i := 2; i*i < n; i++ {
		if isNotPrime[i] {
			continue
		}
		for j := i * i; j < n; j += i {
			isNotPrime[j] = true
		}
	}
	for i := 2; i < n; i++ {
		if !isNotPrime[i] {
			count++
		}

	}
	return count
}

// @lc code=end

