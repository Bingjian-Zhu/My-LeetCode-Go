/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 */

// @lc code=start
func findDuplicate(nums []int) int {
	res:=0
	var existMap = make(map[int]int8, len(nums)-1)
	for i := 0; i < len(nums); i++ {
		existMap[nums[i]]++
	}
	for key, val := range existMap {
		if val > 1 {
			res = key
			break
		}
	}
	return res
}
// @lc code=end

