/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := []int{}
	// preorder(&ans, root)
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return ans
}

func preorder(res *[]int, root *TreeNode) {
	if root != nil {
		*res = append(*res, root.Val)
		preorder(res, root.Left)
		preorder(res, root.Right)
	}
}

// @lc code=end

