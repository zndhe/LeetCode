/*
 * @lc app=leetcode.cn id=889 lang=golang
 *
 * [889] 根据前序和后序遍历构造二叉树
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
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	if len(postorder) == 1 {
		return &TreeNode{postorder[0], nil, nil}
	}
	root := &TreeNode{Val: preorder[0]}
	i := 0
	for ; i < len(postorder); i++ {
		if postorder[i] == preorder[1] {
			break
		}
	}
	leftLen := i + 1
	root.Left = constructFromPrePost(preorder[1:1+leftLen], postorder[:leftLen])
	root.Right = constructFromPrePost(preorder[1+leftLen:], postorder[leftLen:len(postorder)-1])

	return root
}

// @lc code=end

