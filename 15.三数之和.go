/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    var res [][]int
    for i := 0; i < len(nums) - 2; i++ {
        if nums[i] > 0 {
            break
        }
        if i > 0 && nums[i] == nums[i - 1] {
            continue
        }
        left := i + 1
        right := len(nums) - 1
        for left < right {
            sum := nums[i] + nums[left] + nums[right]
            if sum == 0 {
                res = append(res, []int{nums[i], nums[left], nums[right]})
                for left < right && nums[left] == nums[left + 1] {
                    left++
                }
                for left < right && nums[right] == nums[right - 1] {
                    right--
                }
                left++
                right--
            } else if sum < 0 {
                left++
            } else {
                right--
            }
        }
    }
    return res
}
// @lc code=end

