/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 *
 * https://leetcode-cn.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (75.38%)
 * Likes:    1064
 * Dislikes: 0
 * Total Accepted:    132.7K
 * Total Submissions: 175.9K
 * Testcase Example:  '3'
 *
 * 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
 *
 *
 *
 * 示例：
 *
 * 输入：n = 3
 * 输出：[
 * ⁠      "((()))",
 * ⁠      "(()())",
 * ⁠      "(())()",
 * ⁠      "()(())",
 * ⁠      "()()()"
 * ⁠    ]
 *
 *
 */

// @lc code=start

func generate(s string, leftBracketNum int, rightBracketNum int, result *[]string) {
	if leftBracketNum == 0 && rightBracketNum == 0 {
		*result = append(*result, s)
		return
	}

	if leftBracketNum > 0 {
		generate(s+"(", leftBracketNum-1, rightBracketNum, result)
	}

	if rightBracketNum > leftBracketNum {
		generate(s+")", leftBracketNum, rightBracketNum-1, result)
	}
}

func generateParenthesis(n int) []string {
	result := []string{}
	generate("", n, n, &result)
	return result
}

// @lc code=end

