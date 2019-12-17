/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	res := 0
	lenS := len(s)
	if lenS == 0 {
		return res
	}
	strMap := make(map[byte]struct{}, 100)
	left := 0
	right := 0
	for left < lenS && right < lenS {
		if _, ok := strMap[s[right]]; !ok {
			strMap[s[right]] = struct{}{}
			right++
			res = maxInt(res, len(strMap))
		} else {
			delete(strMap, s[left])
			left++
		}
	}
	return res
}

func maxInt(x, y int) int {
	if x < y {
		return y
	}
	return x
}
// @lc code=end

