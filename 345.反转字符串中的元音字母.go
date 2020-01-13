/*
 * @lc app=leetcode.cn id=345 lang=golang
 *
 * [345] 反转字符串中的元音字母
 */

// @lc code=start
// 481/481 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 93.75 % of golang submissions (4 MB)
func reverseVowels(s string) string {
	j := len(s) - 1
	byteArr := []byte(s)
	for i:=0;i<j;i++{
		for j > 0 && byteArr[j]!='a'&& byteArr[j]!='A' && byteArr[j]!='e'&& byteArr[j]!='E' && byteArr[j]!='i'&& byteArr[j]!='I' && byteArr[j]!='o'&& byteArr[j]!='O' && byteArr[j]!='u' && byteArr[j]!='U'{
			j--
		}
		if byteArr[i] =='a'|| byteArr[i] =='A' || byteArr[i] =='e' || byteArr[i] =='E'|| byteArr[i] =='i' || byteArr[i] =='I'|| byteArr[i] =='o' || byteArr[i] =='O'|| byteArr[i] =='u'|| byteArr[i] =='U'{
			byteArr[i],byteArr[j] = byteArr[j],byteArr[i]
			j--
		}
	}
	return string(byteArr[:])
}
// @lc code=end

