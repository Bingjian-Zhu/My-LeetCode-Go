/*
 * @lc app=leetcode.cn id=746 lang=golang
 *
 * [746] 使用最小花费爬楼梯
 */

// @lc code=start
func minCostClimbingStairs(cost []int) int {
	cost = append(cost, 0)
	lenCost := len(cost)
	if lenCost < 2 {
		return 0
	}
	dp := make([]int, lenCost)
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < lenCost; i++ {
		dp[i] = min(cost[i]+dp[i-2], cost[i]+dp[i-1])
	}
	return dp[lenCost-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
// @lc code=end

