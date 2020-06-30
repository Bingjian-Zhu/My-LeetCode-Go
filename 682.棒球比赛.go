/*
 * @lc app=leetcode.cn id=682 lang=golang
 *
 * [682] 棒球比赛
 */

// @lc code=start
func calPoints(ops []string) int {
	sum := 0
	flag := false
	for i, num := range ops {
		if num == "C" {
			for j := i - 1; j >= 0; j-- {
				if ops[j] != "C" {
					sum -= strToInt(ops[j])
					ops[j] = "C"
					break
				}
			}
		} else if num == "+" {
			flag = false
			tmp := 0
			for j := i - 1; j >= 0; j-- {
				if ops[j] != "C" {
					if flag {
						tmp += strToInt(ops[j])
						sum += tmp
						ops[i] = strconv.Itoa(tmp)
						break
					} else {
						tmp += strToInt(ops[j])
						flag = true
					}
				}
			}
		} else if num == "D" {
			for j := i - 1; j >= 0; j-- {
				if ops[j] != "C" {
					tmp := strToInt(ops[j]) * 2
					sum += tmp
					ops[i] = strconv.Itoa(tmp)
					break
				}
			}
		} else {
			sum += strToInt(num)
		}
	}
	return sum
}

func strToInt(str string) int {
	res, _ := strconv.Atoi(str)
	return res
}
// @lc code=end

