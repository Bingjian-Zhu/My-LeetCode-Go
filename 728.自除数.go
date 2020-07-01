/*
 * @lc app=leetcode.cn id=728 lang=golang
 *
 * [728] 自除数
 */

// @lc code=start
func selfDividingNumbers(left int, right int) []int {
	res := []int{}
	for i := left; i <= right; i++ {
		num := i
		for num > 0 {
			tmp := num % 10
			if tmp == 0 || i%tmp != 0 {
				break
			}
			num = num / 10
		}
		if num == 0 {
			res = append(res, i)
		}
	}
	return res
}
// @lc code=end

