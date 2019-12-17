/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
    l := len(nums)
    if l==0{
        return true
    }
    maxv:=nums[0];
    for k,v := range nums{
        if maxv < k{
            return false
        }else if maxv >= l-1{
            return true
        }else if maxv<k+v{
            maxv = k+v
        }
    }
    if(maxv>=l-1){
        return true
    }else{
        return false
    }
    
}
// @lc code=end

