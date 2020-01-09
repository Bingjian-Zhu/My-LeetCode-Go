/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	row:=len(matrix)
	if row == 0{
		return []int{}
	}
	col:=len(matrix[0])
	times := row * col
	res:=make([]int, times)
	l,r,t,b:=0,col-1,0,row-1
	k := 0
	for k<times{
		//左到右
		for i:=l;i<=r && k < times;i++ {
			res[k] = matrix[t][i]
			k++
		}
		t++
		//上到下
		for i:=t;i<=b && k < times;i++{
			res[k] = matrix[i][r]
			k++
		}
		r--
		//右到左
		for i:=r;i>=l && k < times;i--{
			res[k] = matrix[b][i]
			k++
		}
		b--
		//下到上
		for i:=b;i>=t && k < times;i--{
			res[k] = matrix[i][l]
			k++
		}
		l++
	}
	return res
}
// @lc code=end

