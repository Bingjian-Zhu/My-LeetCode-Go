/*
 * @lc app=leetcode.cn id=441 lang=golang
 *
 * [441] 排列硬币
 */

// @lc code=start
// 1336/1336 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 88.89 % of golang submissions (2.2 MB)
func arrangeCoins(n int) int {
	sum := 0
	i := 1
	for sum < n {
		sum += i
		i++
	}
	if sum == n {
		return i-1
	}
	return i - 2
}
// @lc code=end

