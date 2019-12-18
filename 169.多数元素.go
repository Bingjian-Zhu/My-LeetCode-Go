/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	lenNum := len(nums)
	numMap:=make(map[int]int, lenNum/2+1)
	for _,num:=range nums{
		numMap[num]++
		if numMap[num] > lenNum/2{
			return num
		}
	}
	return -1
}
// @lc code=end

