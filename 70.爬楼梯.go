/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	first, second := 1, 1
	for i := 0; i < n; i++ {
		first, second = second, first+second
	}
	return first
}
// @lc code=end

