/*
 * @lc app=leetcode.cn id=212 lang=golang
 *
 * [212] 单词搜索 II
 *
 * https://leetcode-cn.com/problems/word-search-ii/description/
 *
 * algorithms
 * Hard (40.72%)
 * Likes:    186
 * Dislikes: 0
 * Total Accepted:    16K
 * Total Submissions: 38.9K
 * Testcase Example:  '[["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]]\n' +
  '["oath","pea","eat","rain"]'
 *
 * 给定一个二维网格 board 和一个字典中的单词列表 words，找出所有同时在二维网格和字典中出现的单词。
 *
 *
 * 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。
 *
 * 示例:
 *
 * 输入:
 * words = ["oath","pea","eat","rain"] and board =
 * [
 * ⁠ ['o','a','a','n'],
 * ⁠ ['e','t','a','e'],
 * ⁠ ['i','h','k','r'],
 * ⁠ ['i','f','l','v']
 * ]
 *
 * 输出: ["eat","oath"]
 *
 * 说明:
 * 你可以假设所有输入都由小写字母 a-z 组成。
 *
 * 提示:
 *
 *
 * 你需要优化回溯算法以通过更大数据量的测试。你能否早点停止回溯？
 * 如果当前单词不存在于所有单词的前缀中，则可以立即停止回溯。什么样的数据结构可以有效地执行这样的操作？散列表是否可行？为什么？
 * 前缀树如何？如果你想学习如何实现一个基本的前缀树，请先查看这个问题： 实现Trie（前缀树）。
 *
 *
*/

package leetcode

import "fmt"

// @lc code=start
type TrieNode struct {
	word string
	next [26]*TrieNode
}

func dfs(board [][]byte, x int, y int, node *TrieNode, result *[]string) {
	if node == nil {
		return
	}
	fmt.Println(node.word)
	if node.word != "" {
		*result = append(*result, node.word)
		node.word = ""
	}
	originalLetter := board[x][y]
	board[x][y] = '#'

	for _, w := range [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
		newX := x + w[0]
		newY := y + w[1]
		if newX >= 0 && newX < len(board) && newY >= 0 && newY < len(board[0]) && board[newX][newY] != '#' {
			dfs(board, newX, newY, node.next[board[newX][newY]-'a'], result)
		}
	}
	board[x][y] = originalLetter
}

func findWords(board [][]byte, words []string) []string {
	root := &TrieNode{}

	for _, word := range words {
		curr := root
		for i := 0; i < len(word); i++ {
			index := word[i] - 'a'
			next := curr.next[index]
			if next == nil {
				next = &TrieNode{}
			}
			curr.next[index] = next
			curr = next
		}
		curr.word = word
	}

	result := []string{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(board, i, j, root.next[board[i][j]-'a'], &result)
		}
	}
	return result

}

// @lc code=end
