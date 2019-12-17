/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	lenPri:=len(prices)
	if lenPri <= 1{
		return 0
	}
	buy := -prices[0]
	sell :=0
	for i:=1;i<lenPri;i++{
		buy = maxInt(buy,-prices[i])
		sell = maxInt(sell,buy+prices[i])
	}
	return sell
}

func maxInt(x,y int)int{
	if x > y{
		return x
	}
	return y
}
// @lc code=end

