/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	lenS := len(s)
	lenT := len(t)
	if lenS != lenT {
		return false
	}
	sMap := make(map[byte]int, lenS)
	for i := range s {
		sMap[s[i]]++
		sMap[t[i]]--
	}

	for _, val := range sMap {
		if val != 0 {
			return false
		}
	}
	return true
}
// @lc code=end

