/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
func rob(nums []int) int {
    lenNum := len(nums)
	if lenNum == 0{
		return 0
	}else if lenNum == 1{
		return nums[0]
	}
	robFirst := rob1(nums[:lenNum -1])
	robSecond:= rob1(nums[1:])
	if robFirst > robSecond{
		return robFirst
	}
	return robSecond
}

func rob1(nums []int)int{
	lenNum := len(nums)
	if lenNum == 1{
		return nums[0]
	}
	dp := make([]int, lenNum)
	dp[0] = nums[0]
	temp := make([]int,lenNum)
	temp[0] = 0
	if dp[0] < nums[1]{
		dp[1] = nums[1]
		temp[1] = 1
	}else{
		dp[1] = nums[0]
		temp[1] = 0
	}
	for i := 2;i < lenNum;i++{
		if temp[i-1]+1==i{
			if dp[i-1] < nums[i] + dp[i-2]{
				dp[i] = nums[i] + dp[i-2]
				temp[i] = i
			}else{
				dp[i] = dp[i-1]
				temp[i] = temp[i-1]
			}
		}else{
			dp[i] = dp[i-1]+nums[i]
			temp[i] = i
		}
	}
	// fmt.Println(dp)
	// fmt.Println(temp)
	return dp[lenNum-1]
}
// @lc code=end

