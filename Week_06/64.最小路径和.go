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

// @lc code=start
// package leetcode

import "math"

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}

func minPathSumTopDown(grid [][]int, x int, y int, memo map[[2]int]int) int {
	if x >= len(grid) || y >= len(grid[0]) {
		return math.MaxInt64
	}

	if x == len(grid)-1 && y == len(grid[0])-1 {
		return grid[x][y]
	}

	if v, ok := memo[[2]int{x, y}]; ok {
		return v
	}

	n1 := minPathSumTopDown(grid, x+1, y, memo)
	n2 := minPathSumTopDown(grid, x, y+1, memo)
	n := min(n1, n2) + grid[x][y]
	memo[[2]int{x, y}] = n
	return n
}

func minPathSumTD(grid [][]int) int {
	memo := map[[2]int]int{}
	return minPathSumTopDown(grid, 0, 0, memo)
}

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))

	for i := range dp {
		dp[i] = make([]int, len(grid[0]))
	}

	dp[0][0] = grid[0][0]

	for i := 1; i < len(grid); i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < len(grid[0]); i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}

// @lc code=end
