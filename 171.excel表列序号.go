/*
 * @lc app=leetcode.cn id=171 lang=golang
 *
 * [171] Excel表列序号
 */

// @lc code=start
func titleToNumber(s string) int {
	res:=0
	times := 1
    for i := len(s)-1; i >= 0; i-- {
		res+=(int(s[i]-'A') + 1) * times
		times = times * 26
	}
	return res
}
// @lc code=end

