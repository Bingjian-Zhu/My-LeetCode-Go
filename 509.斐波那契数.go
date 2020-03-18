/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

// @lc code=start
// 31/31 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 97.25 % of golang submissions (1.9 MB)
func fib(N int) int {
    first, second := 0, 1
    for i := 0; i < N; i++ {
        first, second = second, first+second
    }
    return first
}
// @lc code=end

