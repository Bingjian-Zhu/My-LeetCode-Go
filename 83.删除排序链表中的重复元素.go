/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
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
	if head == nil {
		return head
	}
	for head.Next != nil {
		last = head
		if head.Val == head.Next.Val {
			//fmt.Printf("last:%d  head:%d\n",last.Val,head.Val)
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			last.Next = head.Next
		} 
		head = head.Next
		if head == nil {
			break
		}
	}
	return res
}
// @lc code=end

