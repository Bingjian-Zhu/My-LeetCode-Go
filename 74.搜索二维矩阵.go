/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	lenM := len(matrix)
    if lenM == 0 ||len(matrix[0])==0 {
		return false
	}
	left, right := 0, lenM-1
	for left <= right {
		mid := (left + right) / 2
		//fmt.Printf("left:%d, right:%d, mid:%d\n", left, right, mid)
		if matrix[mid][0] < target {
			left = mid + 1
		} else if matrix[mid][0] > target {
			right = mid - 1
		} else {
			return true
		}
	}
	row := (left + right) / 2
	//fmt.Printf("left:%d, right:%d, row:%d\n", left, right, row)
	left = 0
	right = len(matrix[row]) - 1
	for left <= right {
		mid := (left + right) / 2
		//fmt.Printf("left:%d, right:%d, mid:%d\n", left, right, mid)
		if matrix[row][mid] < target {
			left = mid + 1
		} else if matrix[row][mid] > target {
			right = mid - 1
		} else {
			return true
		}
	}
	return false
}
// @lc code=end

