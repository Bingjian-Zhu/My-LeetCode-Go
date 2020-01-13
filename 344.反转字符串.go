/*
 * @lc app=leetcode.cn id=344 lang=golang
 *
 * [344] 反转字符串
 */

// @lc code=start
// 478/478 cases passed (40 ms)
// Your runtime beats 97.19 % of golang submissions
// Your memory usage beats 100 % of golang submissions (6.2 MB)
func reverseString(s []byte)  {
	i,j:=0,len(s)-1
	for i<j{
		s[i],s[j] = s[j],s[i]
		i++
		j--
	}
}
// @lc code=end

