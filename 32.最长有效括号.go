/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */

// @lc code=start
func longestValidParentheses(s string) int {
	maxans := 0
	lenS := len(s)
	stack := []int{}
	top := 0 //头指针
	stack = append(stack, -1)
	top++
	for i := 0; i < lenS; i++ {
		if s[i] == '(' {
			stack = append(stack, i)
			top++
		} else {
			top--
			if top == 0 {
				stack = append(stack, i)
				top++
			} else {
				maxans = max(maxans, i-stack[top])
			}
		}
	}
	if maxans%2 != 0 {
		return maxans + 1
	}
	return maxans
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

