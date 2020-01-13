/*
 * @lc app=leetcode.cn id=231 lang=golang
 *
 * [231] 2的幂
 */

// @lc code=start
// 1108/1108 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 94.79 % of golang submissions (2.2 MB)
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
    for n > 1 {
		if n % 2 != 0{
			return false
		}
		tmp := n / 2
		if tmp * 2 != n {
			return false
		}
		n = tmp
	}
	return true
}
// @lc code=end

