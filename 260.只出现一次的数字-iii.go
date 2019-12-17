/*
 * @lc app=leetcode.cn id=260 lang=golang
 *
 * [260] 只出现一次的数字 III
 */

// @lc code=start
func singleNumber(nums []int) []int {
	res:=[]int{}
	var existMap = make(map[int]int8, len(nums)/2+1)
	for i := 0; i < len(nums); i++ {
		existMap[nums[i]]++
	}
	for key, val := range existMap {
		if val < 2 {
			res = append(res,key)
		}
	}
	return res
}
// @lc code=end

