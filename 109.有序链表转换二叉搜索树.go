/*
 * @lc app=leetcode.cn id=109 lang=golang
 *
 * [109] 有序链表转换二叉搜索树
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var slowHead *ListNode

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head != nil && head.Next == nil {
		return &TreeNode{Val: head.Val, Left: nil, Right: nil}
	}
	slowHead = head
	length := getLength(head)
	return buildTree(0, length-1)
}

func buildTree(left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right + 1) / 2
	root := &TreeNode{}
	root.Left = buildTree(left, mid-1)
	root.Val = slowHead.Val
	slowHead = slowHead.Next
	root.Right = buildTree(mid+1, right)
	return root
}
func getLength(head *ListNode) int {
	res := 0
	for ; head != nil; head = head.Next {
		res++
	}
	return res
}

// @lc code=end

