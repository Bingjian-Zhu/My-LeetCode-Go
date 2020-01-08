/*
 * @lc app=leetcode.cn id=392 lang=golang
 *
 * [392] 判断子序列
 */

// @lc code=start
func isSubsequence(s string, t string) bool {
	lenS,lenT := len(s),len(t)
	if lenS > lenT{
		return false
	}
	if s == ""{
		return true
	}
	i:=0
	for j := 0; j < lenT; j++ {
		if s[i] == t[j]{
			i++
			if i == lenS{
				return true
			}
		}
	}
	return false
}
// @lc code=end

