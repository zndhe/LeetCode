/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else {
		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
	}
	return 0
}

//k是到中位数个数，非下标
func getKthElement(nums1, nums2 []int, k int) int {
	//前面index位已经排除
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		newindex1 := min(index1+k/2, len(nums1))
		newindex2 := min(index2+k/2, len(nums2))
		num1, num2 := nums1[newindex1-1], nums2[newindex2-1]
		if num1 < num2 {
			k -= (newindex1 - index1 )
			index1 = newindex1
		} else {
			k -= (newindex2 - index2 )
			index2 = newindex2
		}
	}
	return 0
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

