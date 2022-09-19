/*
 * @lc app=leetcode.cn id=637 lang=golang
 *
 * [637] 二叉树的层平均值
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
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{0}
	}
	ans := []float64{}
	queue := []*TreeNode{root}
	curNum, nextNum, sum, count := 1, 0, 0, 1
	for len(queue) > 0 {
		if curNum > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
				nextNum++
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				nextNum++
			}
			curNum--
			sum += node.Val
		}
		if curNum == 0 {
			ans = append(ans, float64(sum)/float64(count))
			curNum, nextNum, sum, count = nextNum, 0, 0, nextNum
		}
	}
	return ans
}

// @lc code=end

