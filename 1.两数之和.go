/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	index := make(map[int]int)
	var tmp int
	for i, num := range nums {
		tmp = target - num
		if j, ok := index[tmp]; ok {
			return []int{j, i}
		}

		index[num] = i
	}

	return []int{-1, -1}
}
// @lc code=end

