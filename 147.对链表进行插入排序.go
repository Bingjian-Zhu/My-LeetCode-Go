/*
 * @lc app=leetcode.cn id=147 lang=golang
 *
 * [147] 对链表进行插入排序
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func insertionSortList(head *ListNode) *ListNode {
	temp := head
    if temp==nil{
        return temp
    }
	temp = temp.Next
	for temp != nil {
		//log(temp)
		//log(head)
		res := head
		last := head
		isHead := true
		for head != nil {
			if temp.Val <= head.Val {
				tmp := &ListNode{
					Val:  temp.Val,
					Next: nil,
				}
				next := temp.Next
				end := head
				if temp != end {
					for temp != end.Next {
						end = end.Next
					}
				} else {
					head = res
					break
				}
				//fmt.Printf("end:%d\n", end.Val)
				end.Next = next
				tmp.Next = head
				if isHead {
					head = tmp
				} else {
					last.Next = tmp
					head = res
				}
				break
			} else {
				last = head
				//fmt.Printf("last:%d\n", last.Val)
				head = head.Next
				isHead = false
			}
		}
		temp = temp.Next
	}
	return head
}
// @lc code=end

