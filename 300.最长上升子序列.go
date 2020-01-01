/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长上升子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
    for i := 0; i < len(dp); i++ {
        dp[i] = 1
	}
	res:=0
	for i:=0;i<len(nums);i++{
		for j:=0;j<i;j++{
			if nums[j] < nums[i]{
				dp[i] = maxInt(dp[i],dp[j]+1)
			}
		}
		res = maxInt(res,dp[i])
	}
	//fmt.Println(dp)
	return res
}

func maxInt(x,y int) int{
	if x > y {
		return x
	}
	return y
}
// @lc code=end

