/*
 * @lc app=leetcode.cn id=71 lang=golang
 *
 * [71] 简化路径
 */

// @lc code=start
func simplifyPath(path string) string {
	strs := strings.Split(path, "/")
	stack := [1000]string{}
	top := 0
	for i := 0; i < len(strs); i++ {
		if strs[i] == ".." {
			top--
		} else if strs[i] != ""&& strs[i]!="."{
            if top<0{
                top=0
            }
			stack[top]=strs[i]
			top++
		}
	}
	var res bytes.Buffer
	res.WriteString("/")
	for i := 0; i < top; i++ {
		res.WriteString(stack[i])
		if i < top-1 {
			res.WriteString("/")
		}
	}
	return res.String()
}
// @lc code=end

