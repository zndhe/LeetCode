/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
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
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := []int{}

	stack := []*TreeNode{}
	prev := &TreeNode{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 为叶子节点，或已经遍历完了左右子树
		if root.Right == nil || prev == root.Right {
			ans = append(ans, root.Val)
			prev = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}

	}
	return ans
}

// @lc code=end

