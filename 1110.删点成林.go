/*
 * @lc app=leetcode.cn id=1110 lang=golang
 *
 * [1110] 删点成林
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
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	if root == nil {
		return nil
	}
	res, deleteMap := []*TreeNode{}, map[int]bool{}
	for _, v := range to_delete {
		deleteMap[v] = true
	}
	root = dfsDel(root, deleteMap, &res)
	if root != nil {
		res = append(res, root)
	}
	return res
}

//从下往上。确定自己的左右子树
func dfsDel(root *TreeNode, toDel map[int]bool, res *[]*TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	//从下面开始遍历
	root.Left = dfsDel(root.Left, toDel, res)
	root.Right = dfsDel(root.Right, toDel, res)

	if toDel[root.Val] {
		if root.Left != nil {
			*res = append(*res, root.Left)
		}
		if root.Right != nil {
			*res = append(*res, root.Right)
		}
		root = nil
	}
	return root
}

// @lc code=end

