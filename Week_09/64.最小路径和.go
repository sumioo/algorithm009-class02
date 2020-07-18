/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 *
 * https://leetcode-cn.com/problems/minimum-path-sum/description/
 *
 * algorithms
 * Medium (65.53%)
 * Likes:    495
 * Dislikes: 0
 * Total Accepted:    94.7K
 * Total Submissions: 143.8K
 * Testcase Example:  '[[1,3,1],[1,5,1],[4,2,1]]'
 *
 * 给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
 *
 * 说明：每次只能向下或者向右移动一步。
 *
 * 示例:
 *
 * 输入:
 * [
 * [1,3,1],
 * ⁠ [1,5,1],
 * ⁠ [4,2,1]
 * ]
 * 输出: 7
 * 解释: 因为路径 1→3→1→1→1 的总和最小。
 *
 *
 */

package week9

import (
	"fmt"
	"math"
)

// @lc code=start
// package leetcode

// Position (x,y)
type Position [2]int

func minPathSumTopDown(grid [][]int, memo map[Position]int, row int, column int) int {
	if row >= len(grid) || column >= len(grid[0]) {
		return math.MaxInt64
	}
	if row == len(grid)-1 && column == len(grid[0])-1 {
		return grid[row][column]
	}

	position := Position{row, column}
	if v, ok := memo[position]; ok {
		return v
	}

	// down, right := math.MaxInt64,
	// if row+1 < len(grid) && column < len(grid[0]) { //检测能否进入下一层 ps:在终止条件检测好维护和理解点
	// 	down = minPathSumTopDown(grid, memo, row+1, column)
	// }

	// if row < len(grid) && column+1 < len(grid[0]) {
	// 	right = minPathSumTopDown(grid, memo, row, column+1)
	// }
	down := minPathSumTopDown(grid, memo, row+1, column)
	right := minPathSumTopDown(grid, memo, row, column+1)
	pathSum := min(down, right) + grid[row][column]
	memo[position] = pathSum
	return pathSum
}

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	memo := map[Position]int{}
	return minPathSumTopDown(grid, memo, 0, 0)
}

// @lc code=end

func printMinPath(dp [][]int, grid [][]int, row int, column int) {
	if row == 0 && column == 0 {
		fmt.Print(grid[row][column], "->")
		return
	}

	left, up := math.MaxInt64, math.MaxInt64
	if row-1 >= 0 && column >= 0 {
		left = dp[row-1][column]
	}
	if row >= 0 && column-1 >= 0 {
		up = dp[row][column-1]
	}

	if left < up {
		print(dp, grid, row-1, column)
	} else {
		print(dp, grid, row, column-1)
	}
	fmt.Print(grid[row][column], "->")
	return
}
