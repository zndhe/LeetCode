/*
 * @lc app=leetcode.cn id=897 lang=golang
 *
 * [897] 递增顺序搜索树
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
func increasingBST(root *TreeNode) *TreeNode {
	head := &TreeNode{}
	//遍历指针
	temp := head
	buildBST(root, temp)
	return head.Right
}
func buildBST(root, temp *TreeNode) *TreeNode {
	if root == nil {
		return temp
	}
	temp = buildBST(root.Left, temp)
	root.Left = nil
	temp.Right = root
	temp = root
	temp = buildBST(root.Right, temp)
	return temp
}

// @lc code=end

