/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	lenGrid := len(obstacleGrid)
	if lenGrid == 0 ||obstacleGrid[0][0] == 1{
		return 0
	}
	//0表示行不通，与题目0,1互换，方便计算
	//处理第一行，0，1互换，遇到障碍 之后的都设成0
	isZero := false
	for i := 0; i < len(obstacleGrid[0]); i++ {
		if obstacleGrid[0][i] == 0 && !isZero {
			obstacleGrid[0][i] = 1
			continue
		} 
		if obstacleGrid[0][i] == 1 || isZero {
			obstacleGrid[0][i] = 0
			isZero = true
		}
	}
	//处理第一列，0，1互换，遇到障碍 之后的都设成0
	isZero = false
	for i := 1; i < lenGrid; i++ {
		if obstacleGrid[i][0] == 0 && !isZero {
			obstacleGrid[i][0] = 1
			continue
		} 
		if obstacleGrid[i][0] == 1 || isZero  {
			obstacleGrid[i][0] = 0
			isZero = true
		}
	}
	//fmt.Println(obstacleGrid)
	for i := 1; i < lenGrid; i++ {
		for j := 1; j < len(obstacleGrid[i]); j++ {
			//把0，1互换，0表示行不通
			if obstacleGrid[i][j] == 1{
				obstacleGrid[i][j] = 0
			}else{
				obstacleGrid[i][j] = 1
			}
			//dp计算，0不处理
			if obstacleGrid[i][j] == 0 {
				continue
			} else {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			}
		}
	}
	
	return obstacleGrid[lenGrid-1][len(obstacleGrid[0])-1]
}
// @lc code=end

