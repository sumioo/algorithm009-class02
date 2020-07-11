/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 *
 * https://leetcode-cn.com/problems/number-of-islands/description/
 *
 * algorithms
 * Medium (49.47%)
 * Likes:    598
 * Dislikes: 0
 * Total Accepted:    115.2K
 * Total Submissions: 232.6K
 * Testcase Example:  '[["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]'
 *
 * 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
 *
 * 岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。
 *
 * 此外，你可以假设该网格的四条边均被水包围。
 *
 *
 *
 * 示例 1:
 *
 * 输入:
 * 11110
 * 11010
 * 11000
 * 00000
 * 输出: 1
 *
 *
 * 示例 2:
 *
 * 输入:
 * 11000
 * 11000
 * 00100
 * 00011
 * 输出: 3
 * 解释: 每座岛屿只能由水平和/或竖直方向上相邻的陆地连接而成。
 *
 *
 */
 package week4
 
// @lc code=start
func dfs(grid [][]byte, row int, column int) {
	if row >= len(grid) || row < 0 || column >= len(grid[row]) || column < 0 || grid[row][column] != '1' {
		return
	}
	grid[row][column] = '0'
	dfs(grid, row-1, column)
	dfs(grid, row+1, column)
	dfs(grid, row, column-1)
	dfs(grid, row, column+1)
}

func numIslands(grid [][]byte) int {
	count := 0
	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[row]); column++ {
			if grid[row][column] == '1' {
				count++
				dfs(grid, row, column)
			}
		}
	}
	return count
}

// @lc code=end

