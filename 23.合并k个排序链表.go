/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeKLists(lists []*ListNode) *ListNode {
	lenList := len(lists)
	if lenList == 0 {
		return nil
	}
	left := 0
	temp := lenList - 1
	for temp != 0 {
		for right := temp; left < right; right-- {
			lists[left] = mergeTwoLists(lists[left], lists[right])
			left++
		}
		left = 0
		temp = temp / 2
	}
	return lists[0]
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	res := &ListNode{}
	if l1.Val < l2.Val {
		res.Val = l1.Val
		l1 = l1.Next
	} else {
		res.Val = l2.Val
		l2 = l2.Next
	}
	temp := res
	for l1 != nil || l2 != nil {
		if l1 == nil {
			temp.Next = l2
			return res
		}
		if l2 == nil {
			temp.Next = l1
			return res
		}
		if l1.Val < l2.Val {
			temp.Next = l1
			l1 = l1.Next
			temp = temp.Next
		} else {
			temp.Next = l2
			l2 = l2.Next
			temp = temp.Next
		}
	}
	return res
}
// @lc code=end

