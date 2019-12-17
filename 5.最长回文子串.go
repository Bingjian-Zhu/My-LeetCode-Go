/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func preProcess(s string) string {
	n := len(s)
	if n == 0 {
		return "^$"
	}
	ret := "^"
	for i := 0; i < n; i++ {
		ret += "#" + string(s[i])
	}
	ret += "#$"
	return ret

}

// 马拉车算法
func longestPalindrome(s string) string {
	T := preProcess(s)
	n := len(T)
	P := [9000]int{}
	C := 0
	R := 0
	for i := 1; i < n-1; i++ {
		i_mirror := 2*C - i
		if R > i {
			P[i] = minInt(R-i, P[i_mirror]) // 防止超出 R
		} else {
			P[i] = 0 // 等于 R 的情况
		}

		// 碰到之前讲的三种情况时候，需要利用中心扩展法
		for T[i+1+P[i]] == T[i-1-P[i]] {
			P[i]++
		}

		// 判断是否需要更新 R
		if i+P[i] > R {
			C = i
			R = i + P[i]
		}

	}

	// 找出 P 的最大值
	maxLen := 0
	centerIndex := 0
	for i := 1; i < n-1; i++ {
		if P[i] > maxLen {
			maxLen = P[i]
			centerIndex = i
		}
	}
	start := (centerIndex - maxLen) / 2 //最开始讲的求原字符串下标
	return s[start : start+maxLen]
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
// @lc code=end

