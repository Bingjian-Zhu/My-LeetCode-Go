/*
 * @lc app=leetcode.cn id=665 lang=golang
 *
 * [665] 非递减数列
 */

// @lc code=start
func checkPossibility(nums []int) bool {
	isFirst := false
	lenNums := len(nums)
	if lenNums <= 2 {
		return true
	}
	if nums[0] > nums[1] {
		isFirst = true
		nums[0] = nums[1]
	}
	for i := 1; i < lenNums; i++ {
		if nums[i-1] > nums[i] {
			if isFirst {
				return false
			} else {
				isFirst = true
				if nums[i-2] > nums[i] {
					nums[i] = nums[i-1]
				}
			}
		}
	}
	return true
}
// @lc code=end

