/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
	if x == 0 {
		return x
	}
	l, r := 1, x
	res := 0
	for l <= r {
		mid := (l + r) / 2
		if x/mid >= mid {
			l = mid + 1
			res = mid
		} else {
			r = mid - 1
		}
	}
	return res
}

// @lc code=end

