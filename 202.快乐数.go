/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */

// @lc code=start
// 401/401 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 89.41 % of golang submissions (2 MB)
func isHappy(n int) bool {
    for n != 1 {
		tmp:=n
		n = 0
		for tmp > 0 {
			t := tmp % 10
			n += t * t
			tmp /= 10
		}
		if n == 4 {
			return false
		}
	}
	return true
}
// @lc code=end

