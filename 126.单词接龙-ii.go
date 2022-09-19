/*
 * @lc app=leetcode.cn id=126 lang=golang
 *
 * [126] 单词接龙 II
 */

// @lc code=start
func findLadders(beginWord string, endWord string, wordList []string) [][]string {

	preMap := bfs(wordList, beginWord, endWord)
	if preMap == nil {
		return nil
	}

	ans := [][]string{}
	path := []string{}
	dfs(endWord, preMap, &ans, path)
	return ans
}

//从endword进行深度遍历，则不用回溯，应该走的路一定是正确的
//在广度的时候，从beginword开始
//最后能到达endword的一定是最短且完整的路，所以进行深搜即可，不用回溯
//如果是从前面遍历，可能会有一段“死胡同”，故需要回溯
func dfs(endWord string, preMap map[string][]string, ans *[][]string, path []string) {
	path = append([]string{endWord}, path...)
	parents := preMap[endWord]

	if len(parents) == 0 {
		// resPath := make([]string, len(path))
		// copy(resPath, path)
		// path = path[:0]
		path = append(path[:0:0], path...) //append返回的是新的slice
		*ans = append(*ans, path)
		return
	}
	for _, node := range parents {
		dfs(node, preMap, ans, path)
	}
	return
}

func bfs(wordList []string, beginWord, endWord string) map[string][]string {
	preMap := make(map[string][]string, 0) //记录当前单词的父单词
	wordLen := len(wordList)
	//使用过的去除，避免来回循环
	used := make([]bool, wordLen)
	has := false

	for i := range wordList {
		if wordList[i] == endWord {
			has = true
			break
		}

	}
	//将队列中与开头单词一样的去除，不然会产生环
	for i := range wordList {
		if wordList[i] == beginWord {
			used[i] = true //这里很重要
			break
		}
	}
	if !has {
		return nil
	}

	queue := make([]string, 0)
	queue = append(queue, beginWord)

	for len(queue) > 0 {
		qMap := make(map[int]struct{}, 0)

		qLen := len(queue)
		for i := 0; i < qLen; i++ {
			word := queue[i]
			//有路到达ednword，停止，避免后面还有更长的路
			//保证在premap的路是最短的
			if word == endWord {
				return preMap
			}

			for j := range wordList {
				if used[j] == false && hasOneDiff(word, wordList[j]) {
					qMap[j] = struct{}{} //将下一层的index存入，如果单纯是加入队列，可以在这里直接append
					//但是需要在整层完成后才可以赋值false，不然会出现log找到cog将used[cog]=true，然后dog找不到cog了
					preMap[wordList[j]] = append(preMap[wordList[j]], word)
				}
			}
		}
		queue = queue[:0]

		//每一层遍历完，将连接的下一层赋为true，防止下一层直接互相连接
		for k := range qMap {
			queue = append(queue, wordList[k])
			used[k] = true
		}
	}
	return nil
}

func hasOneDiff(a, b string) bool {
	dif := 0

	for i := range a {
		if a[i] != b[i] {
			dif++
			if dif > 1 {
				return false
			}
		}
	}
	return dif == 1
}

// @lc code=end

