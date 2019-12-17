/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	var res = [][]int{}
	var lastSlice = [][]int{}
	lenNums := len(nums)
	if lenNums == 0 {
		return append(res, []int{})
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	lastSlice = append(lastSlice, []int{nums[0]})
	res = append(res, []int{nums[0]})
	for i := 1; i < lenNums; i++ {
		var temp = [][]int{}
		if nums[i-1] != nums[i] {
			lastSlice = res
			temp = append(temp, []int{nums[i]})
		}
		for _, val := range lastSlice {
			copyVal := make([]int, len(val))
			copy(copyVal, val)
			temp = append(temp, append(copyVal, nums[i]))
		}
		res = append(res, temp...)
		lastSlice = temp
	}
	res = append(res, []int{})
	return res
}
// @lc code=end

