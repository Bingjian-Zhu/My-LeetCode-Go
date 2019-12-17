/*
 * @lc app=leetcode.cn id=57 lang=golang
 *
 * [57] 插入区间
 */

// @lc code=start
func insert(intervals [][]int, targe []int) [][]int {
	intervals = append(intervals, targe)
	lenInt := len(intervals)
	for i := 0; i < lenInt; i++ {
		isMerge := false
		for j := 0; j < lenInt; j++ {
			if i != j && mergeCheck(intervals[i], intervals[j]) {
				intervals[i] = mergeTwo(intervals[i], intervals[j])
				intervals = append(intervals[:j], intervals[j+1:]...)
				lenInt--
				isMerge = true
			}
		}
		if isMerge {
			i--
		}
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	return intervals
}

func mergeCheck(x, y []int) bool {
	if x[0] <= y[1] && x[1] >= y[0] {
		return true
	}
	return false
}

func mergeTwo(x, y []int) []int {
	left, right := x[0], x[1]
	if left > y[0] {
		left = y[0]
	}
	if right < y[1] {
		right = y[1]
	}
	return []int{left, right}
}
// @lc code=end

