/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
//使用两个map，时间复杂度O(m+n)
func intersection(nums1 []int, nums2 []int) []int {
	num2Map := make(map[int]struct{},10)
	for i := 0; i < len(nums1); i++ {
		if _, ok := num2Map[nums1[i]]; !ok {
			num2Map[nums1[i]] = struct{}{}
		}
	}
	res := []int{}
	num1Map := make(map[int]struct{},10)
	for i := 0; i < len(nums2); i++ {
		if _, ok := num2Map[nums2[i]]; ok {
			if _, ok := num1Map[nums2[i]]; !ok {
				num1Map[nums2[i]] = struct{}{}
				res = append(res, nums2[i])
			}
		}
	}
	return res
}
// @lc code=end

