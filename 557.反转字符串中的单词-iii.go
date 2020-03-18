/*
 * @lc app=leetcode.cn id=557 lang=golang
 *
 * [557] 反转字符串中的单词 III
 */

// @lc code=start
// 30/30 cases passed (4 ms)
// Your runtime beats 95.91 % of golang submissions
// Your memory usage beats 42.25 % of golang submissions (6.1 MB)
func reverseWords(s string) string {
	var res bytes.Buffer
	arrS := strings.Split(s, " ")
	for i,val :=range arrS{
		if i != len(arrS)-1{
			res.WriteString(reverse(val))
			res.WriteString(" ")
		}else{
			res.WriteString(reverse(val))
		}
	}
	return res.String()
}

func reverse(s string) string{
	b := []byte(s)
	i,j:=0,len(s)-1
	for i < j {
		b[i],b[j] = b[j],b[i]
		i++
		j--
	}
	return string(b)
}
// @lc code=end

