/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	myMap := make(map[byte]struct{}, 9)
	//每行判断
	for i:=0;i<9;i++{
		myMap = make(map[byte]struct{}, 9)
		for j:=0;j<9;j++{
			if board[i][j] >= '0' && board[i][j] <= '9'{
				if _, ok := myMap[board[i][j]]; ok {
					return false
				}
				myMap[board[i][j]] = struct{}{}
			}
		}
		
	}
	//每列判断
	for i:=0;i<9;i++{
		myMap = make(map[byte]struct{}, 9)
		for j:=0;j<9;j++{
			if board[j][i] >= '0' && board[j][i] <= '9'{
				if _, ok := myMap[board[j][i]]; ok {
					return false
				}
				myMap[board[j][i]] = struct{}{}
			}
		}
		
	}
	//每个3x3的矩形判断
	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			myMap := make(map[byte]struct{}, 9)
			//每行判断
			for row:=i*3;row<(i+1)*3;row++{
				for col:=j*3;col<(j+1)*3;col++{
					if board[row][col] >= '0' && board[row][col] <= '9'{
						if _, ok := myMap[board[row][col]]; ok {
							return false
						}
						myMap[board[row][col]] = struct{}{}
					}
				}
			}
		}
	}
	return true
}

// @lc code=end

