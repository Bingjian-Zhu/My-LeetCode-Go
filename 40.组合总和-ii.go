/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return dfs(candidates, target)
}

func dfs(nums []int, target int) [][]int {
	ret := [][]int{}
	for i, n := range nums {
		if i > 0 && nums[i-1] == n {
			continue
		} else if target < n {
			break
		} else if target == n {
			ret = append(ret, []int{n})
			continue
		}
		for _, v := range dfs(nums[i+1:], target-n) {
			ret = append(ret, append([]int{n}, v...))
		}
	}
	return ret
}
// @lc code=end

