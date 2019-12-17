/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */

// @lc code=start
func addBinary(a string, b string) string {
	var temp byte = 0
	lenA := len(a)
	lenB := len(b)
	arrRes := []string{}
	var tmpA, tmpB byte
	var res bytes.Buffer
	for lenA > 0 || lenB > 0 {
		if lenA > 0 {
			lenA--
			tmpA = a[lenA] - '0'
		} else {
			tmpA = 0
		}
		if lenB > 0 {
			lenB--
			tmpB = b[lenB] - '0'
		} else {
			tmpB = 0
		}

		sum := tmpA + tmpB + temp

		if sum > 1 {
			temp = 1
			sum = sum % 2
		} else {
			temp = 0
		}
		arrRes = append(arrRes, string(sum+'0'))
	}
	if temp > 0 {
		arrRes = append(arrRes, string(1+'0'))
	}
	for i := len(arrRes) - 1; i >= 0; i-- {
		res.WriteString(arrRes[i])
	}
	return res.String()
}
// @lc code=end

