/*
 * @lc app=leetcode.cn id=290 lang=golang
 *
 * [290] 单词规律
 */

// @lc code=start
// 33/33 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 35.48 % of golang submissions (2 MB)
func wordPattern(pattern string, str string) bool {
	lenP := len(pattern)
	paMap:=make(map[byte]string, lenP)
	strMap:=make(map[string]struct{}, lenP)
	strSplit := strings.Split(str," ")
	if lenP != len(strSplit){
		return false
	}
	for i:=0;i<lenP;i++{
		if val,ok := paMap[pattern[i]];ok{
			if val != strSplit[i]{
				return false
			}
		}else{
			paMap[pattern[i]] = strSplit[i]
		}
	}
	for _,val := range strSplit{
		if _,ok := strMap[val];!ok{
			strMap[val] = struct{}{}
		}
	}
	if len(strMap) != len(paMap){
		return false
	}
	return true
}
// @lc code=end

