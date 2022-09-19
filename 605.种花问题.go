/*
 * @lc app=leetcode.cn id=605 lang=golang
 *
 * [605] 种花问题
 */

// @lc code=start
func canPlaceFlowers(flowerbed []int, n int) bool {
	m := len(flowerbed)
	count := 0
	pre := -1 //前一位种花的index
	for i := 0; i < m; i++ {
		if flowerbed[i] == 1 {
			if pre < 0 {
				count += i / 2
			} else {
				count += (i - pre - 2) / 2
			}
			if count >= n {
				return true
			}
			pre = i
		}
	}
	if pre < 0 {
		count += (m + 1) / 2
	} else {
		count += (m - pre - 1) / 2
	}
	return count >= n
}

// @lc code=end

