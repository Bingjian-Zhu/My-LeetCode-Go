/*
 * @lc app=leetcode.cn id=412 lang=golang
 *
 * [412] Fizz Buzz
 */

// @lc code=start
// 8/8 cases passed (4 ms)
// Your runtime beats 97.37 % of golang submissions
// Your memory usage beats 100 % of golang submissions (3.4 MB)
func fizzBuzz(n int) []string {
	res:=make([]string, n)
	for i:=1;i<=n;i++{
		if i % 3 == 0 && i % 5 == 0{
			res[i-1] = "FizzBuzz"
		}else if i % 3 == 0{
			res[i-1] = "Fizz"
		}else if i % 5 == 0{
			res[i-1] = "Buzz"
		}else{
			res[i-1] = strconv.Itoa(i)
		}
	}
	return res
}
// @lc code=end

