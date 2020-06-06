/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 *
 * https://leetcode-cn.com/problems/combinations/description/
 *
 * algorithms
 * Medium (73.63%)
 * Likes:    278
 * Dislikes: 0
 * Total Accepted:    53.1K
 * Total Submissions: 72.1K
 * Testcase Example:  '4\n2'
 *
 * 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
 *
 * 示例:
 *
 * 输入: n = 4, k = 2
 * 输出:
 * [
 * ⁠ [2,4],
 * ⁠ [3,4],
 * ⁠ [2,3],
 * ⁠ [1,2],
 * ⁠ [1,3],
 * ⁠ [1,4],
 * ]
 *
 */

/*

 */

// @lc code=start

func backtrace(m int, n int, k int, nums []int, result *[][]int) {
	if len(nums) == k {
		*result = append(*result, append([]int{}, nums...))
		return
	}

	for i := m; i <= n; i++ {
		nums = append(nums, i)
		backtrace(i+1, n, k, nums, result)
		nums = nums[:len(nums)-1]
	}

}

func combine(n int, k int) [][]int {
	result := [][]int{}
	nums := []int{}
	backtrace(1, n, k, nums, &result)
	return result
}

// @lc code=end

