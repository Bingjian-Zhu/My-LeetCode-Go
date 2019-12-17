/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
type arr [][]int

func (t arr) Len() int {
	return len(t)
}

func (t arr) Less(i, j int) bool {
	if t[i][0] == t[j][0] {
		if t[i][1] >= t[j][1] {
			return true
		}
	} else if t[i][0] < t[j][0] {
		return true
	}
	return false
	// if t[i][0] < t[j][0] {
	// 	return true
	// }
	// return false
}

func (t arr) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
func merge(intervals [][]int) [][]int {
	lenInt := len(intervals)
	if lenInt <= 1 {
		return intervals
	}
	sort.Sort(arr(intervals))
	//fmt.Println(intervals)
	index := 1
	for index < lenInt {
		if mergeCheck(intervals[index-1], intervals[index]) {
			intervals[index-1] = mergeTwo(intervals[index-1], intervals[index])
			intervals = append(intervals[:index], intervals[index+1:]...)
			lenInt--
		} else {
			index++
		}

	}
	return intervals
}

func mergeCheck(x, y []int) bool {
	//fmt.Print(x)
	//fmt.Print(y)
	if x[0] <= y[1] && x[1] >= y[0] {
		//fmt.Println("true")
		return true
	}
	//fmt.Println("false")
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

