/*
 * @lc app=leetcode.cn id=485 lang=golang
 *
 * [485] 最大连续1的个数
 */

// @lc code=start
// 41/41 cases passed (44 ms)
// Your runtime beats 98.02 % of golang submissions
// Your memory usage beats 91.38 % of golang submissions (6.4 MB)
func findMaxConsecutiveOnes(nums []int) int {
	res,temp := 0,0
	lenNum:=len(nums)
	for i:=0;i<lenNum;i++{
		if nums[i] == 1{
			temp++
		}
		if nums[i] != 1 || i == lenNum - 1{
			if temp > res {
				res = temp
			}
			temp = 0
		}
	}
	return res
}
// @lc code=end

