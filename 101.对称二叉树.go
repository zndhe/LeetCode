/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
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
 9.19
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}
func isMirror(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}

// @lc code=end

