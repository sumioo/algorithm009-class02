/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N皇后
 *
 * https://leetcode-cn.com/problems/n-queens/description/
 *
 * algorithms
 * Hard (69.42%)
 * Likes:    411
 * Dislikes: 0
 * Total Accepted:    42.8K
 * Total Submissions: 61.3K
 * Testcase Example:  '4'
 *
 * n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
 *
 *
 *
 * 上图为 8 皇后问题的一种解法。
 *
 * 给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。
 *
 * 每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
 *
 * 示例:
 *
 * 输入: 4
 * 输出: [
 * ⁠[".Q..",  // 解法 1
 * ⁠ "...Q",
 * ⁠ "Q...",
 * ⁠ "..Q."],
 *
 * ⁠["..Q.",  // 解法 2
 * ⁠ "Q...",
 * ⁠ "...Q",
 * ⁠ ".Q.."]
 * ]
 * 解释: 4 皇后问题存在两个不同的解法。
 *
 *
 *
 *
 * 提示：
 *
 *
 *
 * 皇后，是国际象棋中的棋子，意味着国王的妻子。皇后只做一件事，那就是“吃子”。当她遇见可以吃的棋子时，就迅速冲上去吃掉棋子。当然，她横、竖、斜都可走一到七步，可进可退。（引用自
 * 百度百科 - 皇后 ）
 *
 *
 */

// @lc code=start

func backtrace(n int, row int, queens []int, result *[][]int) {
	if row == n {
		*result = append(*result, append([]int{}, queens...))
		return
	}
	for column := 0; column < n; column++ {
		if isOk(queens, row, column) {
			queens[row] = column
			backtrace(n, row+1, queens, result)
		}
	}
}

func isOk(queens []int, row int, column int) bool {
	leftUp := column - 1
	rightUp := column + 1
	for i := row - 1; i >= 0; i-- {
		if queens[i] == column || queens[i] == leftUp || queens[i] == rightUp {
			return false
		}
		leftUp--
		rightUp++
	}
	return true
}

func format(n int, result [][]int) [][]string {
	tables := [][]string{}
	dots := make([]byte, n)
	for i := 0; i < n; i++ {
		dots[i] = '.'
	}

	for _, queens := range result {
		table := []string{}
		for _, column := range queens {
			dots[column] = 'Q'
			table = append(table, string(dots))
			dots[column] = '.'
		}
		tables = append(tables, table)
	}
	return tables
}

func solveNQueens(n int) [][]string {
	result := [][]int{}
	backtrace(n, 0, make([]int, n), &result)
	return format(n, result)
}

// @lc code=end

