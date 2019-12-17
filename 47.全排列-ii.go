/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	lenNum := len(nums)
	if lenNum == 0 {
		return [][]int{}
	}
	if lenNum == 1 {
		return [][]int{[]int{nums[0]}}
	}
	res := permuteUnique(nums[0 : lenNum-1])
	return addNum(res, nums[lenNum-1])
}

//在数组中的各个位置插入目标值
func addNum(arr [][]int, target int) [][]int {
	res := [][]int{}
	for k := 0; k < len(arr); k++ {
		lenArrK := len(arr[k])
		for i := 0; i <= lenArrK; i++ {
			tmp := make([]int, lenArrK)
			copy(tmp, arr[k])
			// fmt.Println(tmp)
			temp, _ := addIndex(tmp, i, target)
			//fmt.Println(temp)
			isEqual := false
			for j := 0; j < len(res); j++ {
				if reflect.DeepEqual(res[j], temp) {
					isEqual = true
					break
				}
			}
			if !isEqual {
				res = append(res, temp)
			}
		}
	}
	return res
}

//在数组下标中插入某元素
func addIndex(slice []int, index int, target int) ([]int, error) {
	if index > len(slice) {
		err := errors.New("the index out of slice")
		return slice, err
	}
	tmp := append([]int{}, slice[index:]...)
	slice = append(slice[0:index], target)
	slice = append(slice, tmp...)
	return slice, nil
}
// @lc code=end

