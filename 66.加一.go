/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	temp := 0
	lenDigit := len(digits)
	digits[lenDigit-1]++
	for i := lenDigit - 1; i >= 0; i-- {
		digits[i] = digits[i] + temp
		if digits[i] >= 10 {
			digits[i] = digits[i] % 10
			temp = 1
		} else {
			return digits
		}
	}
	if temp > 0 {
		digits = append([]int{1}, digits[0:]...)
	}
	return digits
}
// @lc code=end

