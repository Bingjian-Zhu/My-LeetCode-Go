/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */

// @lc code=start
// 23/23 cases passed (12 ms)
// Your runtime beats 98.23 % of golang submissions
// Your memory usage beats 58.54 % of golang submissions (6.6 MB)
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

