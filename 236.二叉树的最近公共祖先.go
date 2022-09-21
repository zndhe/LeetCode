/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
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
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	//遇到就返回
	if root.Val == q.Val || root.Val == p.Val {
		return root
	}
	//在左右子树找到p/q的位置
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	//两个都不为nil，说明分开了
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

// @lc code=end

