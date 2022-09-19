/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

 //9.19

// @lc code=start
type Trie struct {
	isWord   bool
	children [26]*Trie
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isWord = true
}

func (this *Trie) Search(word string) bool {
	node := this
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] != nil {
			node = node.children[ch]
			continue
		}
		return false
	}
	return node.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] != nil {
			node = node.children[ch]
			continue
		}
		return false
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end

