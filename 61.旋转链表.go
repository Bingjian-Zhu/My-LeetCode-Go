/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	lenNode := lenListNode(head)
	if k >= lenNode {
		k = k % lenNode
	}
	for k > 0 {
		temp := head
		last := &ListNode{}
		for head.Next != nil {
			last = head
			head = head.Next
		}
		head.Next = temp
		last.Next = nil
		k--
	}
	return head
}

func lenListNode(head *ListNode) int {
	res := 0
	for head != nil {
		res++
		head = head.Next
	}
	return res
}
// @lc code=end

