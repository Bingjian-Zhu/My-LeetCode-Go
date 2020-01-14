/*
 * @lc app=leetcode.cn id=409 lang=golang
 *
 * [409] 最长回文串
 */

// @lc code=start
// 95/95 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 88 % of golang submissions (2.2 MB)
func longestPalindrome(s string) int {
	lenS := len(s)
	strMap := make(map[byte]int, 26)
	for i:=0;i<lenS;i++{
		strMap[s[i]]++
	}
	sum:=0
	isOne:=false
	for _,val := range strMap{
		if val % 2 != 0{
			isOne = true
		}
		sum+=val/2*2
	}
	if isOne {
		return sum + 1
	}
	return sum
}
// @lc code=end

