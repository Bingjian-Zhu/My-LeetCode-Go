/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 */

// @lc code=start
func countPrimes(n int) int {
	sum := 0
	arr := make([]bool, n)
	for i:=2;i<n;i++{
		if !arr[i]{
			sum++
			for j:=i+i;j<n;j=j+i{
				arr[j] = true
			}
		}
	}
	return sum
}

// //暴力法，超时
// func countPrimes(n int) int {
//     if n <= 2 {
// 		return 0
// 	}
// 	sum := 1
// 	for i:=3;i<n;i++{
// 		if i % 2 == 0{
// 			continue
// 		}
// 		if prime(i){
// 			sum++
// 		}
// 	}
// 	return sum
// }

// func prime(n int)bool{
// 	tmp := n / 2
// 	for i:=2;i<=tmp;i++{
// 		if n % i == 0{
// 			return false
// 		}
// 	}
// 	return true
// }
// @lc code=end

