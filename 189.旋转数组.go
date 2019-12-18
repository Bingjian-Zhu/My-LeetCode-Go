/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
func rotate(nums []int, k int)  {
	lenNum := len(nums)
	if k > lenNum {
		k = k % lenNum
	}
	index:=lenNum-k
	temp := append(nums[index:],nums[:index]...)
	for i,num := range temp{
		nums[i] = num
	}
}
// @lc code=end

