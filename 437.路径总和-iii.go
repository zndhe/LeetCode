/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
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
func pathSum(root *TreeNode, targetSum int) int {
	prefixSum := make(map[int64]int)
	prefixSum[0] = 1
	return dfs(root, prefixSum, 0, int64(targetSum))
}
func dfs(root *TreeNode, prefixSum map[int64]int, cur, sum int64) int {
	if root == nil {
		return 0
	}
	cnt := 0
	cur += int64(root.Val)
	cnt += prefixSum[cur-sum]
	prefixSum[cur]++
	cnt += dfs(root.Left, prefixSum, cur, sum)
	cnt += dfs(root.Right, prefixSum, cur, sum)
	prefixSum[cur]--

	return cnt
}

// func pathSum(root *TreeNode, targetSum int) int {
// 	if root == nil {
// 		return 0
// 	}
// 	res := findPath437(root, targetSum)
// 	res += pathSum(root.Left, targetSum)
// 	res += pathSum(root.Right, targetSum)

// 	return res
// }

// //计算包括root结点的路径。(每次都遍历完，相当于n2)
// func findPath437(root *TreeNode, targetSum int) int {
// 	if root == nil {
// 		return 0
// 	}
// 	res := 0
// 	if root.Val == targetSum {
// 		res++
// 	}
// 	res += findPath437(root.Left, targetSum-root.Val)
// 	res += findPath437(root.Right, targetSum-root.Val)
// 	return res
// }

// @lc code=end

