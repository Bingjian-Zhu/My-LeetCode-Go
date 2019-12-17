/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
	lenTri := len(triangle)
	if lenTri == 0 {
		return 0
	} else if lenTri == 1 {
		return triangle[0][0]
	}
	dp := make([][]int,lenTri)
	for i := 0; i < lenTri; i++ {
        dp[i] = make([]int, lenTri)
    }
	dp[0][0] = triangle[0][0]
	for i := 1; i < lenTri; i++ {
		lenTriI := len(triangle[i])
		for j := 0; j < lenTriI; j++ {
			if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j == lenTriI-1 {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = minInt(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
			}
		}
	}
	min := dp[lenTri-1][0]
	for i := 1; i < len(dp[lenTri-1]); i++ {
		if min > dp[lenTri-1][i] {
			min = dp[lenTri-1][i]
		}
	}
	return min
}

func minInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}
// @lc code=end

