/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 */
// @lc code=start
//题解：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/solution/yi-ge-fang-fa-tuan-mie-6-dao-gu-piao-wen-ti-by-l-3/
func maxProfit(prices []int) int {
	n := len(prices)
	dp_i_0 := 0
	dp_i_1 := math.MinInt32
    dp_pre_0 := 0 // 代表 dp[i-2][0]
    for i := 0; i < n; i++ {
        temp := dp_i_0
        dp_i_0 = maxInt(dp_i_0, dp_i_1 + prices[i])
        dp_i_1 = maxInt(dp_i_1, dp_pre_0 - prices[i])
        dp_pre_0 = temp
    }
    return dp_i_0
}

func maxInt(x,y int)int{
	if x > y{
		return x
	}
	return y
}
// @lc code=end

