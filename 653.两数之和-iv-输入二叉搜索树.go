/*
 * @lc app=leetcode.cn id=653 lang=golang
 *
 * [653] 两数之和 IV - 输入二叉搜索树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	nums := make(map[int]int, 0)
	return find(root, k, nums)
}
func find(root *TreeNode, k int, nums map[int]int) bool {
	if root == nil {
		return false
	}
	if _, ok := nums[k-root.Val]; ok {
		return true
	}
	nums[root.Val]++
	return find(root.Left, k, nums) || find(root.Right, k, nums)
}

// @lc code=end

