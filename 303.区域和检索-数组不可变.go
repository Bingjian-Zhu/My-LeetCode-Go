/*
 * @lc app=leetcode.cn id=303 lang=golang
 *
 * [303] 区域和检索 - 数组不可变
 */

// @lc code=start
//DP
type NumArray struct {
    Arr []int
}

func Constructor(nums []int) NumArray {
	sum := make([]int, len(nums))
	for i:=0;i<len(nums);i++{
		if i == 0{
			sum[i] = nums[i]
		}else{
			sum[i] = sum[i-1]+nums[i]
		}
	}
	//fmt.Println(sum)
	return NumArray{ Arr:sum}
}


func (this *NumArray) SumRange(i int, j int) int {
	if i == 0{
		return this.Arr[j]
	}
	return this.Arr[j]-this.Arr[i-1]
}

////常规法
// type NumArray struct {
//     Arr []int
// }


// func Constructor(nums []int) NumArray {
// 	num := NumArray{ Arr:nums}
//     return num
// }


// func (this *NumArray) SumRange(i int, j int) int {
// 	sum:=0
// 	for i <= j{
// 		sum+=this.Arr[i]
// 		i++
// 	}
// 	return sum
// }


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
// @lc code=end

