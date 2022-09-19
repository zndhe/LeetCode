/*
 * @lc app=leetcode.cn id=513 lang=golang
 *
 * [513] 找树左下角的值
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
func findBottomLeftValue(root *TreeNode) int {
	return bfs(root)
	// if root == nil {
	// 	return 0
	// }
	// res, maxHeight := 0, -1
	// dfs(root, 0, &res, &maxHeight)
	// return res
}

func dfs(root *TreeNode, height int, res, maxHeight *int) {
	if height > *maxHeight && root.Left == nil && root.Right == nil {
		*res = root.Val
		*maxHeight = height
	}
	if root.Left != nil {
		dfs(root.Left, height+1, res, maxHeight)
	}
	if root.Right != nil {
		dfs(root.Right, height+1, res, maxHeight)
	}
}

func bfs(root *TreeNode) int {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		next := []*TreeNode{}
		for _, node := range queue {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		if len(next) == 0 {
			return queue[0].Val
		}
		queue = next
	}
	return 0
}

// @lc code=end

