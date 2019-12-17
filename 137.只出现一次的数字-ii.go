/*
 * @lc app=leetcode.cn id=137 lang=golang
 *
 * [137] 只出现一次的数字 II
 */

// @lc code=start
func singleNumber(nums []int) int {
	sort.Ints(nums)
	res:=nums[0]
	for i:=0;i<len(nums);i++{
		if i % 3!=0{
			if nums[i] != res{
				i++
			   break
			}
		}else{
		res = nums[i]
		}
	}
	return res
}
// @lc code=end

