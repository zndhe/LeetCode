/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
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
 //9.19
func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	helper(root, &res)
	return res
}
func helper(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	left := helper(root.Left, res)
	right := helper(root.Right, res)
	*res = max(*res, left+right)
	return max(left, right) + 1

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

