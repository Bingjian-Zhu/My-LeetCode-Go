/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)
	strSum := make(map[int32]int, 16)
	for _, str := range strs {
		var mul int32 = 1
		var sum int32 = 0
		for _, v := range str {
			mul *= v
			sum += v
		}
		sum += mul
		if index, ok := strSum[sum]; ok {
			result[index] = append(result[index], str)
		} else {
			strSum[sum] = len(result)
			result = append(result, []string{str})
		}
	}
	return result
}
// @lc code=end

