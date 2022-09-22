/*
 * @lc app=leetcode.cn id=450 lang=golang
 *
 * [450] 删除二叉搜索树中的节点
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
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else if root.Left == nil {
		return root.Right
	} else if root.Right == nil {
		// root.Left == nil && root.Right == nil 的情况包含在内
		return root.Left
	} else {
		//root.Val == key && root.Left != nil && root.Right != nil
		next, pre := root.Right, root
		for next.Left != nil {
			pre = next
			next = next.Left
		}
		if pre.Val != root.Val {
			pre.Left = next.Right
		} else {
			pre.Right = next.Right
		}
		next.Left = root.Left
		next.Right = root.Right
		return next
	}
	return root
}

// @lc code=end

