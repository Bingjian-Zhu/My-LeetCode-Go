/*
 * @lc app=leetcode.cn id=389 lang=golang
 *
 * [389] 找不同
 */

// @lc code=start
func findTheDifference(s string, t string) byte {
	var res byte = 0
   for i := 0; i < len(s); i++{
	   res ^= s[i]   
   }
   for i := 0 ; i < len(t); i++{
	   res ^= t[i]   
   }
   return res
}
// @lc code=end

