/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 *
 * https://leetcode-cn.com/problems/climbing-stairs/description/
 *
 * algorithms
 * Easy (48.62%)
 * Likes:    1013
 * Dislikes: 0
 * Total Accepted:    197.9K
 * Total Submissions: 406K
 * Testcase Example:  '2'
 *
 * 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
 *
 * 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
 *
 * 注意：给定 n 是一个正整数。
 *
 * 示例 1：
 *
 * 输入： 2
 * 输出： 2
 * 解释： 有两种方法可以爬到楼顶。
 * 1.  1 阶 + 1 阶
 * 2.  2 阶
 *
 * 示例 2：
 *
 * 输入： 3
 * 输出： 3
 * 解释： 有三种方法可以爬到楼顶。
 * 1.  1 阶 + 1 阶 + 1 阶
 * 2.  1 阶 + 2 阶
 * 3.  2 阶 + 1 阶
 *
 *
 */

//记忆化递归
func helper(n int, memo map[int]int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	if v, ok := memo[n]; ok {
		return v
	}

	total := helper(n-1, memo) + helper(n-2, memo)
	memo[n] = total
	return total
}

func climbStairsA(n int) int {
	memo := make(map[int]int, n)
	return helper(n, memo)
}

// @lc code=start
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	twoStepBefore := 1
	oneStepBefore := 2
	total := 0
	for i := 3; i <= n; i++ {
		total = oneStepBefore + twoStepBefore
		twoStepBefore = oneStepBefore
		oneStepBefore = total
	}
	return total
}

// @lc code=end