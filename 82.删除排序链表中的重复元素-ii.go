/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func deleteDuplicates(head *ListNode) *ListNode {
	res := head
	last := head
	isHead := true
	if head == nil {
		return head
	}
	for head.Next != nil {
		if head.Val == head.Next.Val {
			//fmt.Printf("last:%d  head:%d\n",last.Val,head.Val)
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			if isHead {
				res = head.Next
			} else {
				last.Next = head.Next
			}
		} else {
			last = head
			isHead = false
		}
		head = head.Next
		if head == nil {
			break
		}
	}
	return res
}
// @lc code=end

