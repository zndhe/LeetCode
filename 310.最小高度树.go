/*
 * @lc app=leetcode.cn id=310 lang=golang
 *
 * [310] 最小高度树
 */
//9.19
// @lc code=start
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	graph := make([][]int, n)
	degree := make([]int, n)
	//最小高度，转换为找到度在最中间的节点，以他为根节点即最小高度
	for _, v := range edges {
		degree[v[0]]++
		degree[v[1]]++
		graph[v[0]] = append(graph[v[0]], v[1])
		graph[v[1]] = append(graph[v[1]], v[0])
	}
	queue := []int{}
	for i, _ := range graph {
		if degree[i] == 1 {
			queue = append(queue, i)
		}
	}
	return bfs(graph, queue, degree, n)
}

//二分+广搜
func bfs(graph [][]int, queue, degree []int, remain int) []int {
	for remain > 2 {
		s := len(queue)
		remain -= s
		for _, num := range queue {
			for _, v := range graph[num] {
				//结点会被重复读到的问题通过度的减少解决
				//只有当度为1才将它加入下一层
				degree[v]--
				if degree[v] == 1 {
					queue = append(queue, v)
				}
			}
		}
		queue = queue[s:]
	}
	return queue
}

// @lc code=end

