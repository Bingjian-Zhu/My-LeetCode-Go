/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	res, isStart := 0, false
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			res++
			isStart = true
		}
		if s[i] == ' ' && isStart {
			break
		}
	}
	return res
}
// @lc code=end

