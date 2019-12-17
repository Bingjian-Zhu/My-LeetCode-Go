/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	lenNum := len(nums)
	if lenNum == 0 {
		return 0
	}
	sum, res := 0, nums[0]
	for i := 0; i < lenNum; i++ {
		if sum <= 0 {
			sum = nums[i]
			res = max(nums[i], res)
		} else {
			sum += nums[i]
			res = max(sum, res)
		}
	}
	return res
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
// @lc code=end

