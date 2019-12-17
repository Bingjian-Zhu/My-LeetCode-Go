/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
	lenNum := len(nums)
	if lenNum == 0 || (lenNum == 1 && nums[0] >= 0) {
		return 0
	}
	i, res := 0, 1
	for i < lenNum {
		if nums[i]+i+1 >= lenNum {
			return res
		}
		res++
		start := i + 1
		end := i + 1 + nums[i]
		if start > lenNum {
			start = lenNum
		}
		if end > lenNum {
			end = lenNum
		}
		var index, val int
		if start < end {
			index, val = far(nums[start:end])
		}
		if val+i >= lenNum {
			return res
		}
		i = i + index

	}
	return res
}

func far(nums []int) (int, int) {
	lenNum := len(nums)
	index, val := lenNum, nums[lenNum-1]
	//fmt.Printf("%d ", nums[lenNum-1])
	for i := len(nums) - 2; i >= 0; i-- {
		//fmt.Printf("%d ", nums[i])
		if val+index < nums[i]+i+1 {
			val = nums[i]
			index = i + 1
		}
	}
	//fmt.Println()
	return index, val
}
// @lc code=end

