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

package leetcode

// @lc code=start

func backtrace(n int, row int, solution []int, used map[string]map[int]bool, solutions *[][]int) {
	if row == n {
		*solutions = append(*solutions, append([]int{}, solution...))
		return
	}

	for column := 0; column < n; column++ {
		if !used["col"][column] && !used["master"][row+column] && !used["slave"][row-column] {
			solution = append(solution, column)
			used["col"][column] = true
			used["master"][row+column] = true
			used["slave"][row-column] = true
			backtrace(n, row+1, solution, used, solutions)
			solution = solution[:len(solution)-1]
			used["col"][column] = false
			used["master"][row+column] = false
			used["slave"][row-column] = false
		}
	}

}

func format(solution []int) []string {
	dots := make([]byte, len(solution))
	for i := 0; i < len(solution); i++ {
		dots[i] = '.'
	}
	table := []string{}
	for _, column := range solution {
		dots[column] = 'Q'
		table = append(table, string(dots))
		dots[column] = '.'
	}
	return table
}

func solveNQueens(n int) [][]string {
	solutions := [][]int{}
	used := map[string]map[int]bool{
		"col":    map[int]bool{},
		"master": map[int]bool{},
		"slave":  map[int]bool{},
	}
	backtrace(n, 0, []int{}, used, &solutions)
	result := [][]string{}
	for _, solution := range solutions {
		result = append(result, format(solution))
	}
	return result
}

// @lc code=end
