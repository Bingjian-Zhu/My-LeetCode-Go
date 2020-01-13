/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
// 21/21 cases passed (4 ms)
// Your runtime beats 97.84 % of golang submissions
// Your memory usage beats 100 % of golang submissions (3.7 MB)
func moveZeroes(nums []int)  {
	lenN := len(nums)
	if lenN <= 1 {
		return
	}
    for i:=0;i<lenN;i++{
		j := i + 1
		for j < lenN && nums[j] == 0{
			j++
		}
		if j == lenN{
			break
		}
		if nums[i] == 0 {
			nums[i],nums[j] = nums[j],nums[i]
		}
	}
}
// @lc code=end

