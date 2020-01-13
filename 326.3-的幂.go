/*
 * @lc app=leetcode.cn id=326 lang=golang
 *
 * [326] 3的幂
 */

// @lc code=start
// 21038/21038 cases passed (20 ms)
// Your runtime beats 94.67 % of golang submissions
// Your memory usage beats 96.23 % of golang submissions (6.1 MB)
func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
    for n > 1 {
		if n % 3 != 0{
			return false
		}
		tmp := n / 3
		if tmp * 3 != n {
			return false
		}
		n = tmp
	}
	return true
}
// @lc code=end

