/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	n := len(height)
	left,right := 0,n-1
	res := 0
	for left != right{
		s := (right - left) * min(height[left],height[right])
		res = max(res,s)
		if height[left] <= height[right]{
			left++
		}else {
			right--
		}
	}
	return res
}
func max(a int, b int) int{
    if a > b {
        return a
    }else {
        return b
    }
}
func min(a int, b int) int{
    if a < b {
        return a
    }else {
        return b
    }
}
// @lc code=end

