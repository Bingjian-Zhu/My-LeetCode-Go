/*
 * @lc app=leetcode.cn id=172 lang=golang
 *
 * [172] 阶乘后的零
 */

// @lc code=start
func trailingZeroes(n int) int {
	//找规律
	res:=0
    for n > 0 {
		n /= 5
        res += n
    }
	return res
}
// @lc code=end

