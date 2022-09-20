/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	if len(inorder) == 1 {
		return &TreeNode{inorder[0], nil, nil}
	}
	root := &TreeNode{}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == postorder[len(postorder)-1] {
			break
		}
	}
	root.Val = inorder[i]
	root.Left = buildTree(inorder[:i], postorder[:i])
	root.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])

	return root
}

// @lc code=end

