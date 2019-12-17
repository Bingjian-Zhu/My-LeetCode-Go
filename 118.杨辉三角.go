/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */

// @lc code=start
func generate(numRows int) [][]int {
	var res = [][]int{}
	if numRows == 0{
		return res
	}
	res = append(res, []int{1})
	if numRows == 1 {
		return res
	}
	temp := []int{1, 1}
	res = append(res, temp)
	if numRows == 2 {
		return res
	}
	for i := 3; i <= numRows; i++ {
		var target = []int{}
		for k := 0; k < i; k++ {
			if k == 0 || k == i-1 {
				target = append(target, 1)
			} else {
				target = append(target, temp[k-1]+temp[k])
			}
		}
		temp = target
		res = append(res, target)
	}
	return res
}
// @lc code=end

