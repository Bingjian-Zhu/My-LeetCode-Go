/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] 缺失数字
 */

// @lc code=start
func missingNumber(nums []int) int {
    min, max, sum := nums[0], nums[0], nums[0]
    for i := 1; i < len(nums); i++ {
        if min > nums[i] {
            min = nums[i]
        }
        if max < nums[i] {
            max = nums[i]
        }
        sum += nums[i]
    }
    if min > 0 {
        return 0
    }
    if max == len(nums)-1 {
        return max + 1
    }

    return (max+1)*len(nums)/2 - sum
}
// @lc code=end

