/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	max := 0
	i := 0
	j := len(height) - 1
	for i < j {
		tmp := smallInt(height[i], height[j]) * (j - i)
		if tmp > max {
			max = tmp
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}
func smallInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
// @lc code=end

