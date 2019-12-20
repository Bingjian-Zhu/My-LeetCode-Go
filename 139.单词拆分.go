/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	lenS := len(s)
	dp:=make([]bool, lenS + 1)
	dp[0] = true

	for i := 1; i <= lenS; i++ {
		for j := 0; j < i; j++{
			if dp[j] && Contain(wordDict,s[j:i]){
				dp[i] = true
				break
			}
		}
	}
	fmt.Println(dp)
	return dp[lenS]
}

func Contain(slice []string, target string) bool{
	for _,str := range slice{
		if str == target{
			return true
		}
	}
	return false
}
// @lc code=end

