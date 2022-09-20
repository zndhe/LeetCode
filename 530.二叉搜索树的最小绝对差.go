/*
 * @lc app=leetcode.cn id=530 lang=golang
 *
 * [530] 二叉搜索树的最小绝对差
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
func getMinimumDifference(root *TreeNode) int {
	res := math.MaxInt16
	pre := -1
	dfs(root, &res, &pre)
	return res
}

func dfs(root *TreeNode, res, pre *int) {
	if root == nil {
		return
	}
	dfs(root.Left, res, pre)
	if *pre != -1 {
		*res = min(*res, abs(*pre-root.Val))
	}
	*pre = root.Val
	dfs(root.Right, res, pre)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

// @lc code=end

