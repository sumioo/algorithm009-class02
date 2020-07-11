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

/*
重新描述下题目就是：有n个左括号，n个右括号，求可能生成的所有有效括号对

由题目可知我们要在n个左括号n个右括号生成的所有括号对搜索有效的括号对，
可以把搜索过程分为n+n个阶段，每个阶段有两个分叉 1 选择左括号 2 选择右括号
注意括号是有限的，括号没用完时才可以选择
当左右括号都用完时，我们就可以判断生成的括号对是否有效，有效则加入结果集。

剪枝
如果能提前发现生成的括号对是无效的，那么可以提前剪掉这条树枝
或者说有什么条件可以指引我们沿着正确的路径走。


*/
package week4

// @lc code=start

//不剪枝生成所有可能的括号对
func generateAtNoPruning(s string, leftBracketNum int, rightBracketNum int, result *[]string) {
	if leftBracketNum == 0 && rightBracketNum == 0 {
		*result = append(*result, s)
		return
	}

	if leftBracketNum > 0 {
		generate(s+"(", leftBracketNum-1, rightBracketNum, result)
	}

	if rightBracketNum > 0 {
		generate(s+")", leftBracketNum, rightBracketNum-1, result)
	}
}

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
