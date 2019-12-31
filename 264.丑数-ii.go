/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 */

// @lc code=start
func nthUglyNumber(n int) int {
	dp:=[]int{}
	dp = append(dp,1)
	i2,i3,i5:=0,0,0
	for {
		if n == len(dp){
			return dp[n-1]
		}
		tmpI2 := 2*dp[i2]
		tmpI3 := 3*dp[i3]
		tmpI5 := 5*dp[i5]
		if tmpI2 <= tmpI3 && tmpI2 <= tmpI5 {
			if tmpI2 != dp[len(dp)-1]{
				dp = append(dp,tmpI2)
			}
			i2++
		}else if tmpI3 <= tmpI2 && tmpI3 <= tmpI5 {
			if tmpI3 != dp[len(dp)-1]{
				dp = append(dp,tmpI3)
			}
			i3++
		}else if tmpI5 <= tmpI2 && tmpI5 <= tmpI3 {
			if tmpI5 != dp[len(dp)-1]{
				dp = append(dp,tmpI5)
			}
			i5++
		}
	}
	return 0
}
// @lc code=end

