/*
 * @lc app=leetcode.cn id=674 lang=golang
 *
 * [674] 最长连续递增序列
 */

// @lc code=start
func findLengthOfLCIS(nums []int) int {
	max, tmp := 0, 1
	lenNums := len(nums)
	if lenNums == 0 {
		return max
	}
	for i := 1; i < lenNums; i++ {
		if nums[i-1] < nums[i] {
			tmp++
		} else {
			if tmp > max {
				max = tmp
			}
			tmp = 1
		}
	}
	if tmp > max {
		max = tmp
	}
	return max
}
// @lc code=end

