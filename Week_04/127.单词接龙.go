/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 *
 * https://leetcode-cn.com/problems/word-ladder/description/
 *
 * algorithms
 * Medium (41.58%)
 * Likes:    336
 * Dislikes: 0
 * Total Accepted:    43.5K
 * Total Submissions: 102.6K
 * Testcase Example:  '"hit"\n"cog"\n["hot","dot","dog","lot","log","cog"]'
 *
 * 给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord
 * 的最短转换序列的长度。转换需遵循如下规则：
 *
 *
 * 每次转换只能改变一个字母。
 * 转换过程中的中间单词必须是字典中的单词。
 *
 *
 * 说明:
 *
 *
 * 如果不存在这样的转换序列，返回 0。
 * 所有单词具有相同的长度。
 * 所有单词只由小写字母组成。
 * 字典中不存在重复的单词。
 * 你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
 *
 *
 * 示例 1:
 *
 * 输入:
 * beginWord = "hit",
 * endWord = "cog",
 * wordList = ["hot","dot","dog","lot","log","cog"]
 *
 * 输出: 5
 *
 * 解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
 * ⁠    返回它的长度 5。
 *
 *
 * 示例 2:
 *
 * 输入:
 * beginWord = "hit"
 * endWord = "cog"
 * wordList = ["hot","dot","dog","lot","log"]
 *
 * 输出: 0
 *
 * 解释: endWord "cog" 不在字典中，所以无法进行转换。
 *
 */

// @lc code=start

func isOneLetterDiff(wordA string, wordB string) bool {
	n := len(wordA)
	for i := 0; i < n; i++ {
		if wordA[i] != wordB[i] {
			return wordA[i+1:] == wordB[i+1:]
		}
	}
	return false
}

func searchEdges(vertex string, wordList []string) []string {
	edges := []string{}
	for _, word := range wordList {
		if isOneLetterDiff(vertex, word) {
			edges = append(edges, word)
		}
	}
	return edges
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if len(wordList) == 0 {
		return 0
	}
	level := 0
	queue := []string{}
	visited := map[string]bool{}
	queue = append(queue, beginWord)
	visited[beginWord] = true
	found := false

	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			vertex := queue[0]
			queue = queue[1:]

			if vertex == endWord {
				found = true
				goto RETURN
			}

			for _, edge := range searchEdges(vertex, wordList) {
				if !visited[edge] {
					queue = append(queue, edge)
					visited[edge] = true
				}
			}
		}
		level++
	}

RETURN:
	if found {
		return level + 1
	}

	return 0
}

// @lc code=end

