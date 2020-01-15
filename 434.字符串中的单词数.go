/*
 * @lc app=leetcode.cn id=434 lang=golang
 *
 * [434] 字符串中的单词数
 */

// @lc code=start
// 26/26 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 86.67 % of golang submissions (1.9 MB)
func countSegments(s string) int {
	sum := 0
	str := strings.Split(s," ")
	for i:=0;i<len(str);i++{
		if str[i] != ""{
			sum++
		}
	}
	return sum
}
// @lc code=end

