/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 */

// @lc code=start
func maxProfit(prices []int) int {
	lenPri:=len(prices)
	if lenPri <= 1{
		return 0
	}
	fstBuy := -prices[0]
	fstSell:=0
	buy := -prices[0]
	sell :=0
	for i:=1;i<lenPri;i++{
		fstBuy = maxInt(fstBuy, -prices[i]);
        fstSell = maxInt(fstSell, fstBuy + prices[i]);
		buy = maxInt(buy,fstSell-prices[i])
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

