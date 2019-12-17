/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int{
    res := 0
    for _, v := range nums {
        if v != val {
            nums[res] = v
            res++
        }
    }
    return res
}
// @lc code=end

