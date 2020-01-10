/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	lenNum := len(nums)
	numMap := make(map[int]int, lenNum/2)
	for i:=0;i<lenNum;i++{
		if val,ok := numMap[nums[i]];ok{
			if i - val <= k{
				return true
			}
		}
		numMap[nums[i]] = i
	}
	return false
}
// @lc code=end

