/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 翻转字符串里的单词
 */

// @lc code=start
func reverseWords(s string) string {
	var strByte = [][]byte{}
	var temp = []byte{}
	isStart :=false
	var res bytes.Buffer
	for i:=0;i<len(s);i++{
		if s[i]!=' '{
			isStart = true
			temp = append(temp,s[i])
		}else{
			if isStart{
				strByte = append(strByte,temp)
				temp = []byte{}
			}
			isStart = false
		}
	}
	if len(temp)>0{
		strByte = append(strByte,temp)
	}

	for i:=len(strByte)-1;i>0;i--{
		res.WriteString(string(strByte[i]))
		res.WriteString(" ")
	}
	if len(strByte)>0{
	res.WriteString(string(strByte[0]))
	}
	return res.String()
}
// @lc code=end

