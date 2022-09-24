/*
 * @lc app=leetcode.cn id=241 lang=golang
 *
 * [241] 为运算表达式设计优先级
 */

// @lc code=start
func diffWaysToCompute(expression string) []int {
	memo := make(map[string][]int, 0)
	return dfs(expression, memo)
}
func dfs(expression string, memo map[string][]int) []int {
	//出递归条件（如果当前叶子结点为数字，则直接返回即可）
	if ok, val := isNumber(expression); ok {
		return []int{val}
	}
	//memo 用来去掉重复的计算,剪枝
	if _, ok := memo[expression]; ok {
		return memo[expression]
	}
	res := make([]int, 0)
	for i := 0; i < len(expression); i++ {
		ch := expression[i]
		//分治，类似树
		if ch == '+' || ch == '-' || ch == '*' {
			left := dfs(expression[:i], memo)
			right := dfs(expression[i+1:], memo)
			val := 0
			for _, l := range left {
				for _, r := range right {
					switch ch {
					case '+':
						val = l + r
					case '-':
						val = l - r
					case '*':
						val = l * r
					}
					res = append(res, val)
				}
			}
		}
	}
	memo[expression] = res
	return res
}
func isNumber(expression string) (bool, int) {
	val, err := strconv.Atoi(expression)
	if err != nil {
		return false, -1
	}
	return true, val
}

// @lc code=end

