/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] 同构字符串
 */

// @lc code=start
// 30/30 cases passed (4 ms)
// Your runtime beats 59.64 % of golang submissions
// Your memory usage beats 22.92 % of golang submissions (3.3 MB)
func isIsomorphic(s string, t string) bool {
	lenS := len(s)
	sMap := make(map[byte]byte, lenS)
	for i:=0;i<lenS;i++{
		if val,ok := sMap[s[i]];ok{
			if val != t[i]{
				return false
			}
		}else{
			sMap[s[i]]=t[i]
		}
	}
	tMap := make(map[byte]struct{},lenS)
	for i:=0;i<len(t);i++{
		if _,ok := tMap[t[i]];!ok{
			tMap[t[i]] = struct{}{}
		}
	}
	if len(sMap) != len(tMap) {
		return false
	}
	return true
}
// @lc code=end

