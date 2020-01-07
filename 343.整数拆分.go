/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
 */

// @lc code=start
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i:=2; i<=n; i++ {
		for j:=1; j<i; j++ {
			dp[i] = maxInt(dp[i], maxInt(dp[j],j) * (i - j))
		}
	}
	return dp[n]
}

func maxInt(x ,y int) int{
	if x > y{
		return x
	}
	return y
}
// @lc code=end

