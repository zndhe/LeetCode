/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	newHead := &ListNode{Next: head}
	pre, slow, fast := newHead, head, head
	for fast != nil {
		if n <= 0 {
			pre = slow
			slow = slow.Next
		}
		n--
		fast = fast.Next
	}
	pre.Next = slow.Next
	return newHead.Next
}

// @lc code=end

