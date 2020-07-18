/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 *
 * https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/description/
 *
 * algorithms
 * Easy (54.01%)
 * Likes:    1077
 * Dislikes: 0
 * Total Accepted:    240.8K
 * Total Submissions: 441.5K
 * Testcase Example:  '[7,1,5,3,6,4]'
 *
 * 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
 *
 * 如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。
 *
 * 注意：你不能在买入股票前卖出股票。
 *
 *
 *
 * 示例 1:
 *
 * 输入: [7,1,5,3,6,4]
 * 输出: 5
 * 解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
 * ⁠    注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
 *
 *
 * 示例 2:
 *
 * 输入: [7,6,4,3,1]
 * 输出: 0
 * 解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
 *
 *
 */

package leetcode

import "math"

// @lc code=start

/*每个阶段可以选择买入或卖出或休息，如果卖出的话，那么上一个阶段的选择就必须是买入或者休息，所以这里就提示我们必须维护两个状态。
因为当前阶段的某种选择必须依赖于上一阶段的特定选择，也就是说阶段之间的选择是互相制约的，但是如果只定义一个状态的话，会有后效性.
就是说当我在当前阶段选择卖出时我的上一阶段必须是持有股票的。

>无后效性有两层含义，
>第一层含义是，在推导后面阶段的状态的时候，我们只关心前面阶段的状态值，不关心这个状态是怎么一步一步推导出来的。
>第二层含义是，某阶段状态一旦确定，就不受之后阶段的决策影响。

这里定义两个状态 持有股票 和 持有现金
选择买入的话，则状态转移到持有股票，该状态可以由买入操作+上一阶段的持有现金的状态转移而来
选择卖出的话，则状态转移到持有现金，该状态可以由卖出操作+上一阶段的持有股票的状态转移而来
选择休息的话，则状态维持上一阶段的状态，该状态可以由上一阶段的任一状态转移而来

上面的分析适用不限交易次数的时候，搞混了

这题限制只能进行一次交易，
选择买入的话，则状态转移到持有股票，该状态可以由买入操作+初始时持有现金的状态转移而来（即不进行过任何交易）
选择卖出的话，则状态转移到持有现金，该状态可以由卖出操作+上一阶段的持有股票的状态转移而来
选择休息的话，则状态维持上一阶段的状态，该状态可以由上一阶段的对应的状态转移而来

*/

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxProfitA(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	const STOCK = 0
	const CASH = 1
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][STOCK] = -prices[0]
	dp[0][CASH] = 0

	for i := 1; i < len(prices); i++ {
		// dp[i][STOCK] = max(dp[i-1][CASH]-prices[i], dp[i-1][STOCK])
		dp[i][STOCK] = max(-prices[i], dp[i-1][STOCK])             //在买入和休息选择中取最优解
		dp[i][CASH] = max(dp[i-1][STOCK]+prices[i], dp[i-1][CASH]) //在卖出和休息的选择中取最优解
	}
	return dp[len(prices)-1][CASH]

}

func max3(x int, y int, z int) int {
	var temp int
	if x > y {
		temp = x
	} else {
		temp = y
	}

	if temp > z {
		return temp
	}
	return z
}

func maxProfitBackTrace(prices []int, day int, buyChances int, status string, profit int) int {
	//terminator
	if (buyChances == 0 && status == "cash") || day >= len(prices) {
		return profit
	}

	// process current level: make a choice and set state
	// buy
	profitBuy, profitSale, profitRest := math.MinInt64, math.MinInt64, math.MinInt64
	if buyChances > 0 && status == "cash" {
		profitBuy = maxProfitBackTrace(prices, day+1, buyChances-1, "stock", profit-prices[day])
	}

	if status == "stock" {
		profitSale = maxProfitBackTrace(prices, day+1, buyChances, "cash", profit+prices[day])
	}

	profitRest = maxProfitBackTrace(prices, day+1, buyChances, status, profit)

	return max3(profitBuy, profitSale, profitRest)
}

func maxProfit(prices []int) int {
	return maxProfitBackTrace(prices, 0, 1, "cash", 0)
}

// @lc code=end
