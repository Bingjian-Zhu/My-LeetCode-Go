/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 */

// @lc code=start
// 68/68 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 93.33 % of golang submissions (1.9 MB)
func isPerfectSquare(num int) bool {
	tmp := mySqrt(num)
	return tmp * tmp == num
}

func mySqrt(x int) int {
    tmp := math.Sqrt(float64(x))
    return int(tmp)
}
// @lc code=end

