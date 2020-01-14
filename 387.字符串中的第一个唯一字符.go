/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 */

// @lc code=start
// 104/104 cases passed (4 ms)
// Your runtime beats 99.19 % of golang submissions
// Your memory usage beats 83.04 % of golang submissions (5.8 MB)
func firstUniqChar(s string) int {
	lenS := len(s)
	count := make([]int, 26)
    for i := 0; i < lenS; i++ {
        count[s[i] - 'a']++
    }
    for i := 0; i < lenS; i++ {
        if count[s[i] - 'a'] == 1 {
            return i
        }
    }
    return -1
}
// @lc code=end

