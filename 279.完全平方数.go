/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

// @lc code=start
//DP
func numSquares(n int) int {
    if (n < 2) { return 1 }
    dp := make([]int, n + 1)
    for i:=1; i<=n; i++ {
        dp[i] = i
        for j:=1; i>=j*j; j++ {
            if (dp[i] > dp[i - j*j] + 1) {
                dp[i] = dp[i - j*j] + 1
            }
        }
    }
    return dp[n]
}

////找规律
// func numSquares(n int) int {
// 	for n % 4 == 0 {
// 		n /= 4
// 	}
// 	if n % 8 == 7 {
// 		return 4
// 	}
// 	a := 0
// 	for a * a <= n {
// 		b := mySqrt(n - a * a)
// 		if a*a+b*b == n {
// 			if a != 0 && b != 0 {
// 				return 2
// 			} else {
// 				return 1
// 			}
// 		}
// 		a++
// 	}
// 	return 3
// }

// func mySqrt(x int) int {
// 	left, right := 0, x
// 	for left <= right {
// 		mid := (left + right) / 2
// 		if mid*mid > x {
// 			right = mid - 1
// 		} else if mid*mid < x {
// 			left = mid + 1
// 		} else {
// 			return mid
// 		}
// 		if right*right == x {
// 			return right
// 		}
// 	}
// 	return right
// }

// @lc code=end

