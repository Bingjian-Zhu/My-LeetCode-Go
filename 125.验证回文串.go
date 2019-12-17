/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

// @lc code=start
func isPalindrome(s string) bool {
	lenS:=len(s)
	var (
		left = 0
		right = lenS-1
	)
	s = strings.ToLower(s)
	for left<right{
		if !check(s[left]){
			left++
		}else if !check(s[right]){
			right--
		}else{
			if s[left]!=s[right]{
				return false
			}
			left++
			right--
		}
	}
	return true
}

func check(val byte)bool{
	if (val >='a'&& val<='z')||(val>='0' && val<='9'){
		return true
	}
	return false
}
// @lc code=end

