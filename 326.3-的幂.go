/*
 * @lc app=leetcode.cn id=326 lang=golang
 *
 * [326] 3 的幂
 */

// @lc code=start
func isPowerOfThree(n int) bool {
	// for num >= 3 {
	// 	if num%3 == 0 {
	// 		num /= 3
	// 	} else {
	// 		return false
	// 	}
	// }
	// return num == 1
	return n > 0 && (1162261467%n == 0)
}

// @lc code=end

