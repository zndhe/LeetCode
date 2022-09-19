/*
 * @lc app=leetcode.cn id=99 lang=golang
 *
 * [99] 恢复二叉搜索树
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
func recoverTree(root *TreeNode) {
	var pre, fisrt, second *TreeNode
	// pre, fisrt, second = inOrderTraverse(root, pre, fisrt, second)
	stack := []*TreeNode{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre != nil && root.Val < pre.Val {
			second = root
			if fisrt == nil {
				fisrt = pre
			} else {
				break
			}
		}
		pre = root
		root = root.Right
	}
	fisrt.Val, second.Val = second.Val, fisrt.Val
}

func inOrderTraverse(root, pre, fisrt, second *TreeNode) (*TreeNode, *TreeNode, *TreeNode) {
	if root == nil {
		return pre, fisrt, second
	}
	//遍历到最左边开始
	pre, fisrt, second = inOrderTraverse(root.Left, pre, fisrt, second)
	if pre != nil && root.Val < pre.Val {
		second = root
		if fisrt == nil {
			fisrt = pre
		}
	}
	pre = root
	pre, fisrt, second = inOrderTraverse(root.Right, pre, fisrt, second)
	return pre, fisrt, second
}

// @lc code=end

