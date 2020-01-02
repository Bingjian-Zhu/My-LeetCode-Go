/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个有序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	count := len(nums1) + len(nums2)
	//只需遍历两个数组长度和的一半+1
	times := count / 2 + 1
	//记录两个数组的位置
	i,j:=0,0
	//记录最后遍历的两个数
	finalNum,lastNum:=0,0
	for times > 0 {
		lastNum = finalNum
		if i >= len(nums1){
			finalNum = nums2[j]
			j++
		}else if j >=len(nums2){
			finalNum = nums1[i]
			i++
		}else if nums1[i] <= nums2[j] {
			finalNum = nums1[i]
			i++
		}else{
			finalNum = nums2[j]
			j++
		}
		times--
	}
	if count % 2 == 0 {
		return float64(lastNum + finalNum) / 2.0
	}
	return float64(finalNum)
}
// @lc code=end

