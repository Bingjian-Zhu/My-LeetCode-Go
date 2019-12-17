/*
 * @lc app=leetcode.cn id=179 lang=golang
 *
 * [179] 最大数
 */

// @lc code=start
//转成字符串数组进行排序
func largestNumber(nums []int) string {
	var numStr []string
	for _, num := range nums {
		numStr = append(numStr, strconv.Itoa(num))
	}

	sort.Slice(numStr, func(i, j int) bool {
		return numStr[i]+numStr[j] > numStr[j]+numStr[i] 
	})

	var res bytes.Buffer
	isStart := false
	for i := 0; i < len(numStr); i++ {
		if isStart {
			res.WriteString(numStr[i])
		}
		if numStr[i] != "0" && !isStart {
			isStart = true
			res.WriteString(numStr[i])
		}
	}
	tmp := res.String()
	if tmp == "" {
		return "0"
	}
	return tmp
}
// @lc code=end

