/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子序列
 */

// @lc code=start
func maxProduct(nums []int) int {
	max := math.MinInt32
	imax,imin := 1,1
   for i:=0; i<len(nums); i++{
	   if nums[i] < 0{ 
		   imax,imin = imin,imax
	   }
	   imax = maxInt(imax*nums[i], nums[i])
	   imin = minInt(imin*nums[i], nums[i])
	   max = maxInt(max, imax)
   }
   return max
}

func maxInt(x,y int)int{
   if x > y{
	   return x
   }
   return y
}

func minInt(x,y int)int{
   if x < y{
	   return x
   }
   return y
}
// @lc code=end

