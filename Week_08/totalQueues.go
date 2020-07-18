/*
 * @lc app=leetcode.cn id=52 lang=golang
 *
 * [52] N皇后 II
 *
 * https://leetcode-cn.com/problems/n-queens-ii/description/
 *
 * algorithms
 * Hard (78.57%)
 * Likes:    130
 * Dislikes: 0
 * Total Accepted:    26.3K
 * Total Submissions: 33.4K
 * Testcase Example:  '4'
 *
 * n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
 *
 *
 *
 * 上图为 8 皇后问题的一种解法。
 *
 * 给定一个整数 n，返回 n 皇后不同的解决方案的数量。
 *
 * 示例:
 *
 * 输入: 4
 * 输出: 2
 * 解释: 4 皇后问题存在如下两个不同的解法。
 * [
 * [".Q..",  // 解法 1
 * "...Q",
 * "Q...",
 * "..Q."],
 *
 * ["..Q.",  // 解法 2
 * "Q...",
 * "...Q",
 * ".Q.."]
 * ]
 *
 *
 *
 *
 * 提示：
 *
 *
 *
 * 皇后，是国际象棋中的棋子，意味着国王的妻子。皇后只做一件事，那就是“吃子”。当她遇见可以吃的棋子时，就迅速冲上去吃掉棋子。当然，她横、竖、斜都可走一或七步，可进可退。（引用自
 * 百度百科 - 皇后 ）
 *
 *
 */

package week8

/*
假如位置在第k层 位置n被皇后占据，对于下一层来说，那么列 /，\位置n,n-1,n+1都是不可用的，
对于下二层来说n,n-2,n+2都是不可用的，那么在每一层的操作时我们不仅需要标记当前的选择对于下一层的不可用位置，
还要扩散之前层的标记。
我们可以用一维数组来标记皇后的可用位置。有点滚动复用的思想，但对于
数组来说元素的移动时间复杂度太高了，但位运算就很适合这样的操作。
*/
// @lc code=start
func totalNQueens(n int) int {
	if n < 1 {
		return 0
	}
	count := 0
	backtrace(n, 0, uint32(0), uint32(0), uint32(0), &count)
	return count
}

func backtrace(n int, row int, cols uint32, pie uint32, na uint32, count *int) {
	if row >= n {
		(*count)++
	}
	bits := (^(cols | pie | na)) & ((uint32(1) << n) - 1) //获取当前所有空位
	for bits != 0 {
		p := bits & (-bits)      //取最低位1
		bits = bits & (bits - 1) //清除最低位1 表示在p位置放入皇后，
		backtrace(n, row+1, cols|p, (pie|p)<<1, (na|p)>>1, count)
	}
}

// @lc code=end
