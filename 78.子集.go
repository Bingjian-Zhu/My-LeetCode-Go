/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	var res = [][]int{}
	for _, num := range nums {
		for _, val := range res {
			copyVal := make([]int, len(val))
			copy(copyVal, val)
			temp := append(copyVal, num)
			res = append(res, temp)
		}
		res = append(res, []int{num})
	}
	res = append(res, []int{})
	return res
}
// @lc code=end

