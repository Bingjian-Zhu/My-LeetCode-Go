/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
    var (
		i = 0
		j = len(numbers)-1
	)
	for i<j{
		if target - numbers[i]>numbers[j]{
			i++
		}else if target - numbers[i]<numbers[j]{
			j--
		}else{
			return []int{i+1,j+1}
		}
	}
	return []int{i+1,j+1}
}
// @lc code=end

