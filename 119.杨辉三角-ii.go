/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 */

// @lc code=start
func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	temp := []int{1, 1}
	if rowIndex == 1 {
		return temp
	}
	for i := 3; i <= rowIndex+1; i++ {
		var target = []int{}
		for k := 0; k < i; k++ {
			if k == 0 || k == i-1 {
				target = append(target, 1)
			} else {
				target = append(target, temp[k-1]+temp[k])
			}
		}
		temp = target
	}
	return temp
}
// @lc code=end

