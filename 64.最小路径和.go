/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	//dp算法，dp[i][j] = min(dp[i][j]+dp[i-1][j], dp[i][j]+dp[i][j-1])
    for i:=0; i < len(grid); i++{
		for j :=0; j<len(grid[0]);j++{
			if i == 0 {
				if j == 0{
					continue
				}
				grid[i][j] = grid[i][j] +grid[i][j-1]
			}else if j == 0{
				grid[i][j] = grid[i][j] + grid[i-1][j]
			}else{
				grid[i][j] = min(grid[i][j] + grid[i-1][j],grid[i][j] + grid[i][j-1])
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func min(x,y int)int{
	if x < y {
		return x
	}
	return y
}
// @lc code=end

