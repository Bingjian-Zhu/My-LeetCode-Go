/*
 * @lc app=leetcode.cn id=168 lang=golang
 *
 * [168] Excel表列名称
 */

// @lc code=start
func convertToTitle(n int) string {
	var res bytes.Buffer
	temp:=[]string{}
	for n > 0{
			n--
			temp = append(temp,string(n%26 + 'A'))
			n = n/26
	}
	for i:=len(temp)-1;i>=0;i--{
		res.WriteString(temp[i])
	}
	return res.String()
}
// @lc code=end

