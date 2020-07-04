/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 *
 * https://leetcode-cn.com/problems/sudoku-solver/description/
 *
 * algorithms
 * Hard (61.40%)
 * Likes:    457
 * Dislikes: 0
 * Total Accepted:    32.3K
 * Total Submissions: 52.1K
 * Testcase Example:  '[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]'
 *
 * 编写一个程序，通过已填充的空格来解决数独问题。
 *
 * 一个数独的解法需遵循如下规则：
 *
 *
 * 数字 1-9 在每一行只能出现一次。
 * 数字 1-9 在每一列只能出现一次。
 * 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
 *
 *
 * 空白格用 '.' 表示。
 *
 *
 *
 * 一个数独。
 *
 *
 *
 * 答案被标成红色。
 *
 * Note:
 *
 *
 * 给定的数独序列只包含数字 1-9 和字符 '.' 。
 * 你可以假设给定的数独只有唯一解。
 * 给定数独永远是 9x9 形式的。
 *
 *
 */

package leetcode

// @lc code=start

func backtrace(board [][]byte, n int, used map[string]map[int]map[byte]bool) bool {
	if n == 81 {
		return true
	}

	row := n / 9
	column := n % 9

	if board[row][column] != '.' {
		return backtrace(board, n+1, used)
	} else {
		for i := byte('1'); i <= '9'; i++ {
			subIndex := (row/3)*3 + column/3
			if !used["row"][row][i] && !used["col"][column][i] && !used["sub"][subIndex][i] {
				used["row"][row][i] = true
				used["col"][column][i] = true
				used["sub"][subIndex][i] = true
				board[row][column] = i
				if backtrace(board, n+1, used) {
					return true
				}
				used["row"][row][i] = false
				used["col"][column][i] = false
				used["sub"][subIndex][i] = false
			}
		}
		board[row][column] = '.'
		return false
	}
}

func solveSudoku(board [][]byte) {
	used := map[string]map[int]map[byte]bool{
		"row": map[int]map[byte]bool{},
		"col": map[int]map[byte]bool{},
		"sub": map[int]map[byte]bool{},
	}

	for i := 0; i < 9; i++ {
		used["row"][i] = map[byte]bool{}
		used["col"][i] = map[byte]bool{}
		used["sub"][i] = map[byte]bool{}
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			subIndex := (i/3)*3 + j/3
			if board[i][j] != '.' {
				used["row"][i][board[i][j]] = true
				used["col"][j][board[i][j]] = true
				used["sub"][subIndex][board[i][j]] = true
			}
		}
	}

	backtrace(board, 0, used)
}

// @lc code=end
