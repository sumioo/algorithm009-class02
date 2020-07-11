/*
 * @lc app=leetcode.cn id=126 lang=golang
 *
 * [126] 单词接龙 II
 *
 * https://leetcode-cn.com/problems/word-ladder-ii/description/
 *
 * algorithms
 * Hard (31.60%)
 * Likes:    264
 * Dislikes: 0
 * Total Accepted:    19K
 * Total Submissions: 49.5K
 * Testcase Example:  '"hit"\n"cog"\n["hot","dot","dog","lot","log","cog"]'
 *
 * 给定两个单词（beginWord 和 endWord）和一个字典 wordList，找出所有从 beginWord 到 endWord
 * 的最短转换序列。转换需遵循如下规则：
 *
 *
 * 每次转换只能改变一个字母。
 * 转换后得到的单词必须是字典中的单词。
 *
 *
 * 说明:
 *
 *
 * 如果不存在这样的转换序列，返回一个空列表。
 * 所有单词具有相同的长度。
 * 所有单词只由小写字母组成。
 * 字典中不存在重复的单词。
 * 你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
 *
 *
 * 示例 1:
 *
 * 输入:
 * beginWord = "hit",
 * endWord = "cog",
 * wordList = ["hot","dot","dog","lot","log","cog"]
 *
 * 输出:
 * [
 * ⁠ ["hit","hot","dot","dog","cog"],
 * ["hit","hot","lot","log","cog"]
 * ]
 *
 *
 * 示例 2:
 *
 * 输入:
 * beginWord = "hit"
 * endWord = "cog"
 * wordList = ["hot","dot","dog","lot","log"]
 *
 * 输出: []
 *
 * 解释: endWord "cog" 不在字典中，所以不存在符合要求的转换序列。
 *
 */

// @lc code=start

/*
解题关键
1. 如何获取边
2. 如何得到到达某个点的全部最短路径

*/
package week4

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

func reverse(words []string) {
	i, j := 0, len(words)-1
	for i < j {
		words[i], words[j] = words[j], words[i]
		i++
		j--
	}
}

func backtrace(nextWord string, words []string, edgeTo map[string][]string, result *[][]string) {
	if edgeTo[nextWord] == nil {
		*result = append(*result, append([]string{}, words...))
		reverse((*result)[len(*result)-1])
		return
	}

	for _, word := range edgeTo[nextWord] {
		words = append(words, word)
		backtrace(word, words, edgeTo, result)
		words = words[0 : len(words)-1]
	}
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	if len(wordList) == 0 {
		return nil
	}
	queue := []string{}
	visited := map[string]bool{}
	edgeTo := map[string][]string{}
	result := [][]string{}
	queue = append(queue, beginWord)
	visited[beginWord] = true
	found := false

	for len(queue) > 0 && !found {
		n := len(queue)
		nextLevelVertexs := []string{}
		for i := 0; i < n; i++ {
			vertex := queue[0]
			queue = queue[1:]
			for _, edge := range searchEdges(vertex, wordList) {
				if !visited[edge] {
					if edge == endWord {
						found = true
					}
					edgeTo[edge] = append(edgeTo[edge], vertex)
					nextLevelVertexs = append(nextLevelVertexs, edge)
				}
			}

		}
		for _, vertex := range nextLevelVertexs {
			if !visited[vertex] {
				queue = append(queue, vertex)
			}
			visited[vertex] = true
		}
	}

	if found {
		backtrace(endWord, []string{endWord}, edgeTo, &result)
	}
	return result

}

// @lc code=end
