/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	arr:= make([][]int, n)
	for i:=0;i<n;i++{
		arr[i]=make([]int,n)
	}
	num:=1
	l,r,t,b:=0,n-1,0,n-1
	target := n*n
	for num <= target{
		//左到右
		for i:=l;i<=r;i++ {
			arr[t][i] = num
			num++
		}
		t++
		//上到下
		for i:=t;i<=b;i++{
			arr[i][r] = num
			num++
		}
		r--
		//右到左
		for i:=r;i>=l;i--{
			arr[b][i] = num
			num++
		}
		b--
		//下到上
		for i:=b;i>=t;i--{
			arr[i][l] = num
			num++
		}
		l++
	}
	return arr
}

// @lc code=end

