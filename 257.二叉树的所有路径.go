/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
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
func binaryTreePaths(root *TreeNode) []string {
	ans := []string{}
	dfs(root, "", &ans)
	return ans
}

func dfs(root *TreeNode, path string, ans *[]string) {
	if root != nil {
		path += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			*ans = append(*ans, path)
			return
		}
		path += "->"
		dfs(root.Left, path, ans)
		dfs(root.Right, path, ans)
	}
}

// @lc code=end

