/*
 * @lc app=leetcode.cn id=538 lang=golang
 *
 * [538] 把二叉搜索树转换为累加树
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
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	head := root
	sum := 0
	// dfs(root, &sum)
	stack := []*TreeNode{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Right
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		sum += root.Val
		root.Val = sum
		root = root.Left
	}
	return head
}
func dfs(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	dfs(root.Right, sum)
	root.Val += *sum
	*sum = root.Val
	dfs(root.Left, sum)
}

// @lc code=end

