/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	} else {
		num := 0
		tmp := x
		for x > 0 {
			num = num*10 + x%10
			x = x / 10
		}
		return tmp == num 
	}
}
// @lc code=end

