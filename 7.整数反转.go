/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	isNegative := false
	if x < 0 {
		isNegative = true
		x = x * -1
	}
	res := 0.0
	minVal := -math.Pow(2, 31)
	maxVal := math.Pow(2, 31) - 1
	for x > 0 {
		k := x % 10

		res = res*10.0 + float64(k)
		if res < minVal || res > maxVal {
			return 0
		}
		x = x / 10

	}
	if isNegative {
		return int(res * -1)
	}
	return int(res)
}
// @lc code=end

