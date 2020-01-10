/*
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] 存在重复元素
 */

// @lc code=start
// 18/18 cases passed (16 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 99.59 % of golang submissions (6 MB)
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

