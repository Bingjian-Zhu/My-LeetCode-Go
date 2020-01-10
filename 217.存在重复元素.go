/*
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] 存在重复元素
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	lenNum := len(nums)
	numMap := make(map[int]struct{}, lenNum)
	for i:=0;i<lenNum;i++{
		if _,ok := numMap[nums[i]];ok{
			return true
		}
		numMap[nums[i]] = struct{}{}
	}
	return false
}
// @lc code=end

