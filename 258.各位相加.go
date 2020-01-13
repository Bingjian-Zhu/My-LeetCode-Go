/*
 * @lc app=leetcode.cn id=258 lang=golang
 *
 * [258] 各位相加
 */

// @lc code=start
// 1101/1101 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 97.83 % of golang submissions (2.2 MB)
func addDigits(num int) int {
    return (num - 1) % 9 + 1
}
//循环
// func addDigits(num int) int {
//     for num > 9 {
// 		tmp:=num
// 		num = 0
// 		for tmp > 0 {
// 			num += tmp % 10
// 			tmp /= 10
// 		}
// 	}
// 	return num
// }
// @lc code=end

