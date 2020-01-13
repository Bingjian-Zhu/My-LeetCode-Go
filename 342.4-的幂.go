/*
 * @lc app=leetcode.cn id=342 lang=golang
 *
 * [342] 4的幂
 */

// @lc code=start
// 1060/1060 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (2.1 MB)
func isPowerOfFour(num int) bool {
	if num <= 0 {
		return false
	}
    for num > 1 {
		if num % 4 != 0{
			return false
		}
		tmp := num / 4
		if tmp * 4 != num {
			return false
		}
		num = tmp
	}
	return true
}
// @lc code=end

