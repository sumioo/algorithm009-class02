/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 *
 * https://leetcode-cn.com/problems/valid-perfect-square/description/
 *
 * algorithms
 * Easy (43.37%)
 * Likes:    129
 * Dislikes: 0
 * Total Accepted:    34K
 * Total Submissions: 78.5K
 * Testcase Example:  '16'
 *
 * 给定一个正整数 num，编写一个函数，如果 num 是一个完全平方数，则返回 True，否则返回 False。
 *
 * 说明：不要使用任何内置的库函数，如  sqrt。
 *
 * 示例 1：
 *
 * 输入：16
 * 输出：True
 *
 * 示例 2：
 *
 * 输入：14
 * 输出：False
 *
 *
 */
 package week4
 
// @lc code=start
func isPerfectSquare(num int) bool {
	left, right := 0, num

	for left <= right {
		mid := left + (right-left)/2
		x := mid * mid
		if x == num {
			return true
		} else if x < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

// @lc code=end

