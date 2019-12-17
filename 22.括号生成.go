/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
    res := make([]string,0)
	backTrack(&res, "", 0, 0, n)
	return res
}

func backTrack(res *[]string, s string, open int, close int, max int) {
	if len(s) == max*2 {
		*res = append(*res, s)
		return
	}
	if open < max {
		backTrack(res, s+"(", open+1, close, max)
	}
	if close < open {
		backTrack(res, s+")", open, close+1, max)
	}
}
// @lc code=end

