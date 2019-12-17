/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 报数
 */

// @lc code=start
func countAndSay(n int) string {
	result := "1"
	if n == 1 {
		return result
	}

	for i := 2; i <= n; i++ {
		count := 0
		tmp := ""
		pre := result[0]
		for id := range result {
			if pre != result[id] {
				tmp += strconv.Itoa(count) + string(pre)
				pre = result[id]
				count = 1
				continue
			}
			count++
		}
		tmp += strconv.Itoa(count) + string(pre)
		result = tmp
	}

	return result
}
// @lc code=end

