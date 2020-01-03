/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
   dp:=make([]int, amount + 1)
   for i:=1;i<amount + 1;i++{
	   dp[i] = amount + 1
   }
   for i:= 1;i<=amount;i++{
	   for j:=0;j<len(coins);j++{
		   if i - coins[j] >=0{
			dp[i] = minInt(dp[i],dp[i-coins[j]]+1)
		   }
	   }
   }
   if dp[amount] == amount + 1{
	   return -1
   }
   return dp[amount]
}

func minInt(x,y int) int{
	if x < y{
		return x
	}
	return y
}
// @lc code=end

