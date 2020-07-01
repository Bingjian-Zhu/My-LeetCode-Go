/*
 * @lc app=leetcode.cn id=693 lang=golang
 *
 * [693] 交替位二进制数
 */

// @lc code=start
func hasAlternatingBits(n int) bool {
	str := convertToBin(n)
	if len(str) <= 1 {
		return true
	}
	for i := 1; i < len(str); i++ {
		if str[i-1] == str[i] {
			return false
		}
	}
	return true
}

func convertToBin(n int) string {
	var b string
	switch {
	case n == 0:
		b += "0"
	case n > 0:
		//strcov.Itoa 将 1 转为 "1" , string(1)直接转为assic码
		for ; n > 0; n /= 2 {
			b = strconv.Itoa(n%2) + b
		}
	case n < 0:
		n = n * -1
		// fmt.Println("变为整数：",n)
		s := convertToBin(n)
		//取反
		for i := 0; i < len(s); i++ {
			if s[i:i+1] == "1" {
				b += "0"
			} else {
				b += "1"
			}
		}
		//转化为整形，之后加1 这里必须要64，否则在转化过程中可能会超出范围
		n, err := strconv.ParseInt(b, 2, 64)
		if err != nil {
			fmt.Println(err)
		}
		//转为bin
		//+1
		b = convertToBin(int(n + 1))
	}
	return b
}

// @lc code=end

