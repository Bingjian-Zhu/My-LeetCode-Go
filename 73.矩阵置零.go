/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 */

// @lc code=start
func setZeroes(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				setZero(matrix, i, j)
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == -11 {
				matrix[i][j] = 0
			}
		}
	}
}

func setZero(matrix [][]int, row, col int) {
	for i := 0; i < len(matrix); i++ {
        if matrix[i][col]!=0{
        matrix[i][col] = -11
        }
	}
	for i := 0; i < len(matrix[0]); i++ {
        if matrix[row][i]!=0{
		matrix[row][i] = -11
        }
	}  
}
// @lc code=end

