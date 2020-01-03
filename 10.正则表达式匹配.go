/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

// @lc code=start
func isMatch(s string, p string) bool {
    m, n := len(p)+1, len(s)+1
	dp := make([][]bool, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]bool, n)
	}
	
	//初始值
	dp[0][0] = true
	for i:=2;i<m;i++ {
		if i % 2 == 0 && p[i-1] == '*'{
			dp[i][0] = dp[i-2][0]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if p[i-1] == s[j-1] || p[i-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[i-1] == '*' {
				if i > 1 { //防止脚标越界
					if p[i-2] == s[j-1] || p[i-2] == '.' {
						dp[i][j] = dp[i][j-1] || dp[i-2][j-1] || dp[i-2][j]
					} else {
						dp[i][j] = dp[i-2][j]
					}
				}
			}
		}
	}
	return dp[m-1][n-1]
}
// @lc code=end

