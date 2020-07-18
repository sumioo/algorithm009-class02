/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 *
 * https://leetcode-cn.com/problems/unique-paths-ii/description/
 *
 * algorithms
 * Medium (33.00%)
 * Likes:    380
 * Dislikes: 0
 * Total Accepted:    87.1K
 * Total Submissions: 240.9K
 * Testcase Example:  '[[0,0,0],[0,1,0],[0,0,0]]'
 *
 * 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
 *
 * 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
 *
 * 现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
 *
 *
 *
 * 网格中的障碍物和空位置分别用 1 和 0 来表示。
 *
 * 说明：m 和 n 的值均不超过 100。
 *
 * 示例 1:
 *
 * 输入:
 * [
 * [0,0,0],
 * [0,1,0],
 * [0,0,0]
 * ]
 * 输出: 2
 * 解释:
 * 3x3 网格的正中间有一个障碍物。
 * 从左上角到右下角一共有 2 条不同的路径：
 * 1. 向右 -> 向右 -> 向下 -> 向下
 * 2. 向下 -> 向下 -> 向右 -> 向右
 *
 *
 */

package leetcode

// @lc code=start

/*
f(x,y) = f(x-1,y) + f(x,y-1) if (x-1,y) == 0 && f(x,y-1) == 0
f(x,y) = f(x-1,y) if (x-1,y) == 0
f(x,y) = f(x,y-1) if (x,y-1) == 0
*/

type Position63 [2]int

func uniquePathsWithObstaclesTopDown(obstacleGrid [][]int, memo map[Position63]int, row int, column int) int {
	if row >= len(obstacleGrid) || column >= len(obstacleGrid[0]) || obstacleGrid[row][column] == 1 {
		return 0
	}

	if row == len(obstacleGrid)-1 && column == len(obstacleGrid[0])-1 {
		return 1
	}

	position := Position63{row, column}
	if v, ok := memo[position]; ok {
		return v
	}

	pathNums := (uniquePathsWithObstaclesTopDown(obstacleGrid, memo, row+1, column) +
		uniquePathsWithObstaclesTopDown(obstacleGrid, memo, row, column+1))
	memo[position] = pathNums
	return pathNums
}

func uniquePathsWithObstaclesA(obstacleGrid [][]int) int {
	memo := map[Position63]int{}
	return uniquePathsWithObstaclesTopDown(obstacleGrid, memo, 0, 0)
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}
	rowCount, columnCount := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, rowCount)
	for i := range dp {
		dp[i] = make([]int, columnCount)
	}
	if obstacleGrid[0][0] == 0 {
		dp[0][0] = 1
	}

	for row := 0; row < rowCount; row++ {
		for column := 0; column < columnCount; column++ {
			if obstacleGrid[row][column] == 1 || (row == 0 && column == 0) {
				continue
			}
			if row-1 >= 0 && obstacleGrid[row-1][column] != 1 {
				dp[row][column] += dp[row-1][column]
			}
			if column-1 >= 0 && obstacleGrid[row][column-1] != 1 {
				dp[row][column] += dp[row][column-1]
			}
		}
	}

	return dp[rowCount-1][columnCount-1]
}

// @lc code=end
