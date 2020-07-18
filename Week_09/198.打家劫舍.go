/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 *
 * https://leetcode-cn.com/problems/house-robber/description/
 *
 * algorithms
 * Easy (44.63%)
 * Likes:    938
 * Dislikes: 0
 * Total Accepted:    151.4K
 * Total Submissions: 329.6K
 * Testcase Example:  '[1,2,3,1]'
 *
 *
 * 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
 *
 * 给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
 *
 *
 *
 * 示例 1：
 *
 * 输入：[1,2,3,1]
 * 输出：4
 * 解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
 * 偷窃到的最高金额 = 1 + 3 = 4 。
 *
 * 示例 2：
 *
 * 输入：[2,7,9,3,1]
 * 输出：12
 * 解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
 * 偷窃到的最高金额 = 2 + 9 + 1 = 12 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= nums.length <= 100
 * 0 <= nums[i] <= 400
 *
 *
 */

package leetcode

import "math"

// @lc code=start

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func robBackTrace(nums []int, n int, preChoice string) int {
	if n == 0 {
		if preChoice == "rob" {
			return 0
		}
		return nums[0]
	}

	//process current choice: make a choice and set state then drill down

	r1, r2 := math.MinInt64, math.MinInt64
	if preChoice != "rob" {
		r1 = robBackTrace(nums, n-1, "rob") + nums[n]
	}
	r2 = robBackTrace(nums, n-1, "unrob")

	return max(r1, r2)

}

func robBackTraceB(nums []int, n int) int {
	if n == 0 {
		return nums[0]
	}

	if n == 1 {
		return max(nums[0], nums[1])
	}

	return max(robBackTraceB(nums, n-1), robBackTraceB(nums, n-2)+nums[n])
}

func rob(nums []int) int {
	// return robBackTrace(nums, len(nums)-1, "")
	return robBackTraceB(nums, len(nums)-1)

}

// @lc code=end
