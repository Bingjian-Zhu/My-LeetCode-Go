/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 */

// @lc code=start
//题解：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/solution/yi-ge-fang-fa-tuan-mie-6-dao-gu-piao-wen-ti-by-l-3/
func maxProfit(k int, prices []int) int {
    n := len(prices)
    if k > n / 2{
        return maxProfit_k_inf(prices)
	}
	dp := make([][][]int, n)
	for i:=0;i<n;i++{
		dp[i] = make([][]int,k+1)
		for j :=0;j<len(dp[i]);j++{
			dp[i][j] = make([]int,2)
		}
	}
    //dp := new int[n][k + 1][2];
    for i := 0; i < n; i++{
        for k := k; k >= 1; k-- {
            if i - 1 == -1 { 
                /* 处理 base case */
                dp[i][k][0] = 0
                dp[i][k][1] = -prices[i]
                continue
            }
            dp[i][k][0] = maxInt(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
            dp[i][k][1] = maxInt(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])   
		}
	}
    return dp[n - 1][k][0]
}

func maxProfit_k_inf(prices []int) int{
    n := len(prices)
	dp_i_0 := 0
	dp_i_1 := math.MinInt32
    for i := 0; i < n; i++ {
        temp := dp_i_0
        dp_i_0 = maxInt(dp_i_0, dp_i_1 + prices[i])
        dp_i_1 = maxInt(dp_i_1, temp - prices[i])
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

