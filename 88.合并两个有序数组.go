/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	pos := m + n - 1
	m--
	n--
	for m >=0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[pos] = nums1[m]
			m--
		} else {
			nums1[pos] = nums2[n]
			n--
		}
		pos--
	}
	for n >= 0 {
		nums1[pos] = nums2[n]
		n--
		pos--
	}
}

// @lc code=end

