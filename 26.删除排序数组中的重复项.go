/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	lenNum := len(nums)
	if lenNum <= 1 {
		return lenNum
	}
	temp := 0
	res := 1
	for i := 1; i < lenNum; i++ {
		if nums[temp] != nums[i] {
			temp = i
			nums[res] = nums[i]
			res++
		}
	}
	return res
}
// @lc code=end

