/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res = &ListNode{}
	cur := res
	var count int = 0
	for l1 != nil || l2 != nil || count != 0 {
		if l1 != nil {
			cur.Val = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			cur.Val += l2.Val
			l2 = l2.Next
		}
		cur.Val += count
		if cur.Val >= 10 {
			cur.Val = cur.Val % 10
			count = 1
		} else {
			count = 0
		}
		if l1 == nil && l2 == nil && count == 0 {
			return res
		}
		cur.Next = &ListNode{}
		cur = cur.Next
	}

	return res
}
// @lc code=end

