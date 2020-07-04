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

package leetcode

// @lc code=start
func dfs(grid [][]byte, x int, y int, visited [][]int) {
	visited[x][y] = 1
	for _, w := range [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
		newX := x + w[0]
		newY := y + w[1]
		if newX >= 0 && newX < len(grid) && newY >= 0 && newY < len(grid[0]) && grid[newX][newY] == '1' && visited[newX][newY] == 0 {
			dfs(grid, newX, newY, visited)
		}
	}
}

//归约为查找连通分量
func numIslandsA(grid [][]byte) int {
	visited := make([][]int, len(grid))
	for i := range visited {
		visited[i] = make([]int, len(grid[0]))
	}
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' && visited[i][j] == 0 {
				dfs(grid, i, j, visited)
				count++
			}
		}
	}
	return count
}

type UF struct {
	count int
	id    []int
}

func newUF(n int) *UF {
	x := make([]int, n)
	for i := range x {
		x[i] = i
	}
	return &UF{id: x, count: n}
}

func (u *UF) find(p int) int {
	root := p
	for root != u.id[root] {
		root = u.id[root]
	}

	for p != u.id[p] {
		p, u.id[p] = u.id[p], root
	}

	return root
}

func (u *UF) union(p int, q int) {
	r1 := u.find(p)
	r2 := u.find(q)
	if r1 == r2 {
		return
	}
	u.id[r1] = r2
	u.count--
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	row := len(grid)
	col := len(grid[0])
	length := row * col
	uf := newUF(length + 1)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				for _, w := range [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
					x := i + w[0]
					y := j + w[1]
					if x >= 0 && x < row && y >= 0 && y < col && grid[x][y] == '1' {
						uf.union(i*col+j, x*col+y)
					}
				}
			} else {
				uf.union(i*col+j, length)
			}
		}
	}
	return uf.count - 1
}

// @lc code=end
