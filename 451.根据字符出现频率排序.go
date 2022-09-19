/*
 * @lc app=leetcode.cn id=451 lang=golang
 *
 * [451] 根据字符出现频率排序
 */

// @lc code=start
func frequencySort(s string) string {
	cnt := map[byte]int{}
	for i := range s {
		cnt[s[i]]++
	}
	type bucket struct {
		ch  byte
		cnt int
	}
	buckets := make([]bucket, 0, len(cnt))
	for k, v := range cnt {
		buckets = append(buckets, bucket{k, v})
	}

	//然后对切片进行排序
	sort.Slice(buckets, func(i, j int) bool {
		return buckets[i].cnt > buckets[j].cnt
	})

	res := []byte{}
	for _, v := range buckets {
		res = append(res, bytes.Repeat([]byte{v.ch}, v.cnt)...)
	}
	return string(res)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

