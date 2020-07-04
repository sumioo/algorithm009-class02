/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 *
 * https://leetcode-cn.com/problems/implement-trie-prefix-tree/description/
 *
 * algorithms
 * Medium (67.04%)
 * Likes:    330
 * Dislikes: 0
 * Total Accepted:    42K
 * Total Submissions: 62.1K
 * Testcase Example:  '["Trie","insert","search","search","startsWith","insert","search"]\n' +
  '[[],["apple"],["apple"],["app"],["app"],["app"],["app"]]'
 *
 * 实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。
 *
 * 示例:
 *
 * Trie trie = new Trie();
 *
 * trie.insert("apple");
 * trie.search("apple");   // 返回 true
 * trie.search("app");     // 返回 false
 * trie.startsWith("app"); // 返回 true
 * trie.insert("app");
 * trie.search("app");     // 返回 true
 *
 * 说明:
 *
 *
 * 你可以假设所有的输入都是由小写字母 a-z 构成的。
 * 保证所有输入均为非空字符串。
 *
 *
*/

package leetcode

// @lc code=start

type TrieNode struct {
	isWord bool
	next   [26]*TrieNode
}

type Trie struct {
	root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{root: &TrieNode{}}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	curr := this.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		next := curr.next[index]
		if next == nil {
			next = &TrieNode{}
		}
		curr.next[index] = next
		curr = next
	}
	curr.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.getPrefix(word)
	if node == nil {
		return false
	}
	return node.isWord
}

func (this *Trie) getPrefix(prefix string) *TrieNode {
	curr := this.root
	for i := 0; i < len(prefix); i++ {
		index := prefix[i] - 'a'
		next := curr.next[index]
		if next == nil {
			return nil
		}
		curr = next
	}
	return curr
}

func hasWord(root *TrieNode) bool {
	if root == nil {
		return false
	}

	if root.isWord {
		return true
	}

	for _, node := range root.next {
		if hasWord(node) {
			return true
		}
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this.getPrefix(prefix)
	return hasWord(node)
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end
