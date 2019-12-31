/*
 * @lc app=leetcode.cn id=263 lang=golang
 *
 * [263] 丑数
 */

// @lc code=start
func isUgly(num int) bool {
	if num == 1{
		return true
	}
	if num <= 0{
		return false
	}
	lastNum := num
    for num > 1 {
		if num % 2 == 0{
			num /= 2
		}
		if num % 3 == 0{
			num /= 3
		}
		if num % 5 == 0{
			num /= 5
		}
		if lastNum == num{
			return false
		}
		lastNum = num
	}
	return true
}
// @lc code=end

