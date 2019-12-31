/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] 最大正方形
 */

// @lc code=start
func maximalSquare(matrix [][]byte) int {
	lenMat:=len(matrix)
	dp := make([][]int, lenMat)
	for i:=0;i<lenMat;i++{
		dp[i] = make([]int,len(matrix[i]))
	}
	max:=0
	for i:=0;i<lenMat;i++{
		for j:=0;j<len(matrix[i]);j++{
			if i == 0 || j == 0 {
				max = maxTwo(max,int(matrix[i][j])-'0')
				dp[i][j] = int(matrix[i][j])-'0'
			}else{
				if matrix[i][j] == '1'{
					dp[i][j] = minThree(dp[i-1][j],dp[i-1][j-1],dp[i][j-1])+1
					max = maxTwo(max,dp[i][j])
				}else{
					dp[i][j]= 0
				}
			}
		}
	}
	//fmt.Println(dp)
	return max*max
}

func maxTwo(x,y int)int{
	if x > y{
		return x
	}
	return y
}

func minThree(x,y,z int)int{
	if x <=y && x <= z{
		return x
	}else if y <=z && y <= x{
		return y
	}else {
		return z
	}
}
// @lc code=end

