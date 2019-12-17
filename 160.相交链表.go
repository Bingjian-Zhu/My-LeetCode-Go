/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func getIntersectionNode(headA, headB *ListNode) *ListNode {
	mapList := make(map[*ListNode]struct{},20)
	for headA != nil{
		mapList[headA] = struct{}{}
		headA = headA.Next
	}
	for headB != nil{
		if _, ok := mapList[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
// @lc code=end

