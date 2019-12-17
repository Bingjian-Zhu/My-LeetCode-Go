/*
 * @lc app=leetcode.cn id=80 lang=golang
 *
 * [80] 删除排序数组中的重复项 II
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	lenNum := len(nums)
	if lenNum <= 1 {
		return lenNum
	}
	l := 1
	isTwo := false
	for i := 1; i < lenNum; i++ {
		if nums[i-1] != nums[i] {
			isTwo = false
			nums[l] = nums[i]
			l++
		} else if !isTwo && nums[i-1] == nums[i] {
			isTwo = true
			nums[l] = nums[i]
			l++
		}
	}
	return l
}
// @lc code=end

