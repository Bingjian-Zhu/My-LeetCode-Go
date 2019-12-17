/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */

// @lc code=start
func firstMissingPositive(nums []int) int {
	lenNum := len(nums)

	for i := 0; i < lenNum; i++ {
		for nums[i] > 0 && nums[i] <= lenNum && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i := 0; i < lenNum; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return lenNum + 1
}
// @lc code=end

