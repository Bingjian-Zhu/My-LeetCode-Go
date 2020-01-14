/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */

// @lc code=start
// 315/315 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 73.68 % of golang submissions (2.7 MB)
func addStrings(num1 string, num2 string) string {
	var res bytes.Buffer
	sum,tmp := 0,0
	lenNum1 := len(num1)
	lenNum2 := len(num2)
	if lenNum1 < lenNum2 {
		num1,num2 = num2,num1
		lenNum1,lenNum2 = lenNum2,lenNum1
	}
	j := lenNum2 - 1
	for i:=lenNum1-1; i>=0; i-- {
		if j < 0 {
			sum = int(num1[i]-'0') + tmp
		}else{
			sum = int((num1[i]-'0') + (num2[j]-'0')) + tmp
		}
		tmp = sum / 10
		res.WriteString(strconv.Itoa(sum % 10))
		j--
	}
	if tmp != 0 {
		res.WriteString("1")
	}
	return reverseString(res.String())
}

func reverseString(s string) string {
    runes := []rune(s)

    for i, j := 0, len(runes)-1; i < j; i, j = i + 1, j - 1 {
        runes[i], runes[j] = runes[j], runes[i] 
    } 

    return string(runes)
}
// @lc code=end

