/*
 * @lc app=leetcode.cn id=304 lang=golang
 *
 * [304] 二维区域和检索 - 矩阵不可变
 */

// @lc code=start
type NumMatrix struct {
	Arr [][]int
}

//缓存行
func Constructor(matrix [][]int) NumMatrix {
	lenMat := len(matrix)
    sum := make([][]int, lenMat)
	for i:=0;i<lenMat;i++{
		lenRow := len(matrix[i])
		sum[i] = make([]int,lenRow)
		for j := 0;j < lenRow;j++{
			if j == 0{
				sum[i][j] = matrix[i][j]
			}else{
				sum[i][j] = sum[i][j-1]+matrix[i][j]
			}
		}
	}
	//fmt.Println(sum)
	return NumMatrix{ Arr:sum}
}

//每行计算
func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum:=0
	for i := row1;i<=row2;i++{
		if col1 == 0{
			sum += this.Arr[i][col2]
		}else{
			sum+= this.Arr[i][col2]-this.Arr[i][col1-1]
		}
	}
	return sum
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
// @lc code=end

