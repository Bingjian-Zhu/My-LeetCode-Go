/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 */

// @lc code=start
// 34/34 cases passed (52 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (7.1 MB)
func findDisappearedNumbers(nums []int) []int {
	res := []int{}
	lenNum := len(nums)
	for i:=0;i<lenNum;i++{
		nums[abs(nums[i])-1] = -abs(nums[abs(nums[i])-1])
	}
	for i:=0;i<lenNum;i++{
		if nums[i] > 0 {
			res = append(res,i+1)
		}
	}
	return res
}

func abs(x int)int{
	if x < 0 {
		return -x
	}
	return x
}
//map
// func findDisappearedNumbers(nums []int) []int {
// 	res := []int{}
// 	lenNum := len(nums)
// 	numMap := make(map[int]struct{}, lenNum)
// 	for i:=0;i<lenNum;i++{
// 		numMap[nums[i]] = struct{}{}
// 	}
// 	for i:=1;i<=lenNum;i++{
// 		if _,ok := numMap[i];!ok{
// 			res = append(res,i)
// 		}
// 	}
// 	return res
// }
// @lc code=end

