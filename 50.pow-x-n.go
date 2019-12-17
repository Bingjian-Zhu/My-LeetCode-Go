/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
    if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	var abs_n int
	var res float64
	if n < 0 {
		abs_n = 0 - n
	} else {
		abs_n = n
	}
	temp := myPow(x, abs_n / 2)
	if abs_n % 2 == 0 {
		res = temp * temp
	} else {
		res = x * temp * temp
	}
	if n < 0 {
		res = 1 / res
	}
	return res
}
// @lc code=end

