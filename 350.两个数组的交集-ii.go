/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */

// @lc code=start
//一个map，分别记录数字和出现的次数，当nums2出现这个数字时，次数减一，当次数为0时，删除这个key
func intersect(nums1 []int, nums2 []int) []int {
	numMap := make(map[int]int,10)
	for i := 0; i < len(nums1); i++ {
		if _, ok := numMap[nums1[i]]; !ok {
			numMap[nums1[i]] = 1
		} else {
			numMap[nums1[i]]++
		}
	}
	res := []int{}
	for i := 0; i < len(nums2); i++ {
		if _, ok := numMap[nums2[i]]; ok {
			res = append(res, nums2[i])
			numMap[nums2[i]]--
			if numMap[nums2[i]] == 0 {
				delete(numMap, nums2[i])
			}
		}
	}
	return res
}
// @lc code=end

