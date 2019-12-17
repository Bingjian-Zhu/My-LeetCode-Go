/*
 * @lc app=leetcode.cn id=168 lang=golang
 *
 * [168] Excel表列名称
 */

// @lc code=start
func convertToTitle(n int) string {
	var res bytes.Buffer
	arr:=[]string{"A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z"}
	temp:=[]string{}
	for n > 0{
			n--
			temp = append(temp,arr[n%26])
			n = n/26
	}
	for i:=len(temp)-1;i>=0;i--{
		res.WriteString(temp[i])
	}
	return res.String()
}
// @lc code=end

