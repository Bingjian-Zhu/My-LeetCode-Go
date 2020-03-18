/*
 * @lc app=leetcode.cn id=506 lang=golang
 *
 * [506] 相对名次
 */

// @lc code=start
// 17/17 cases passed (12 ms)
// Your runtime beats 92.11 % of golang submissions
// Your memory usage beats 57.14 % of golang submissions (6 MB)
func findRelativeRanks(nums []int) []string {
	res:=[]string{}
	temp := make([]int, len(nums))
	copy(temp,nums)
	sort.Slice(temp, func(i, j int) bool {
		return temp[i] > temp[j] 
	})
	rankMap := make(map[int]string)
	for index,val := range temp{
		if index == 0{
			rankMap[val] = "Gold Medal"
		}else if index == 1{
			rankMap[val] = "Silver Medal"
		}else if index == 2{
			rankMap[val] = "Bronze Medal"
		}else{
			rankMap[val] = strconv.Itoa(index+1)
		}
	}
	for _,val := range nums{
		res = append(res,rankMap[val])
	}
	return res
}
// @lc code=end

