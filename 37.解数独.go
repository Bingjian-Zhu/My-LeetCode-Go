/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 */

// @lc code=start
func solveSudoku(board [][]byte)  {
    // 三个布尔数组 表明 行, 列, 还有 3*3 的方格的数字是否被使用过
	rowUsed := make([][]bool, 9)
	for i:=0;i<9;i++{
		rowUsed[i] = make([]bool,10)
	}
	colUsed:= make([][]bool, 9)
	for i:=0;i<9;i++{
		colUsed[i] = make([]bool,10)
	}
	boxUsed:=make([][][]bool,3)
	for i:=0;i<3;i++{
		boxUsed[i] = make([][]bool,3)
		for j:=0;j<3;j++{
			boxUsed[i][j] = make([]bool,10)
		}
	}
	// 初始化
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			num := board[row][col] - '0'
			if 1 <= num && num <= 9 {
				rowUsed[row][num] = true
				colUsed[col][num] = true
				boxUsed[row/3][col/3][num] = true
			}
		}
	}
	// 递归尝试填充数组 
	recusiveSolveSudoku(board, rowUsed, colUsed, boxUsed, 0, 0)
}

func recusiveSolveSudoku(board [][]byte, rowUsed [][]bool, colUsed [][]bool, boxUsed[][][]bool, row int, col int) bool {
	// 边界校验, 如果已经填充完成, 返回true, 表示一切结束
	if col == len(board[0]) {
		col = 0
		row++
		if row == len(board) {
			return true
		}
	}
	// 是空则尝试填充, 否则跳过继续尝试填充下一个位置
	if board[row][col] == '.' {
		// 尝试填充1~9
		for num := 1; num <= 9; num++ {
			canUsed := !(rowUsed[row][num] || colUsed[col][num] || boxUsed[row/3][col/3][num])
			if canUsed {
				rowUsed[row][num] = true
				colUsed[col][num] = true
				boxUsed[row/3][col/3][num] = true
				
				board[row][col] = (byte)('0' + num)
				if recusiveSolveSudoku(board, rowUsed, colUsed, boxUsed, row, col + 1){
					return true
				}
				board[row][col] = '.'
				
				rowUsed[row][num] = false
				colUsed[col][num] = false
				boxUsed[row/3][col/3][num] = false
			}
		}
	} else {
		return recusiveSolveSudoku(board, rowUsed, colUsed, boxUsed, row, col + 1)
	}
	return false
}

// @lc code=end

