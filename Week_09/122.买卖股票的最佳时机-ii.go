/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 *
 * https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/description/
 *
 * algorithms
 * Easy (59.83%)
 * Likes:    779
 * Dislikes: 0
 * Total Accepted:    184K
 * Total Submissions: 299.4K
 * Testcase Example:  '[7,1,5,3,6,4]'
 *
 * 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
 *
 * 设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
 *
 * 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
 *
 *
 *
 * 示例 1:
 *
 * 输入: [7,1,5,3,6,4]
 * 输出: 7
 * 解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
 * 随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。
 *
 *
 * 示例 2:
 *
 * 输入: [1,2,3,4,5]
 * 输出: 4
 * 解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4
 * 。
 * 注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
 * 因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
 *
 *
 * 示例 3:
 *
 * 输入: [7,6,4,3,1]
 * 输出: 0
 * 解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= prices.length <= 3 * 10 ^ 4
 * 0 <= prices[i] <= 10 ^ 4
 *
 *
 */

package leetcode

import (
	"fmt"
	"math"
)

// @lc code=start
/*
第一维度是天数，第二维度持股或现金
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
	dp := make([][2]int, len(prices))
	const STOCK = 0
	const CASH = 1
	dp[0][STOCK] = -prices[0]
	dp[0][CASH] = 0

	for i := 1; i < len(prices); i++ {
		dp[i][STOCK] = max(dp[i-1][STOCK], dp[i-1][CASH]-prices[i])
		dp[i][CASH] = max(dp[i-1][CASH], dp[i-1][STOCK]+prices[i])
	}
	return dp[len(prices)-1][CASH]
}

// 没有重复子问题,只能通过剪枝优化，不过好像也没什么可以减枝的
func maxProfitBackTrace2(prices []int, day int, buyChances int, status string, profit int) int {
	//terminator
	if (buyChances == 0 && status == "cash") || day >= len(prices) {
		return profit
	}

	// process current level: make a choice and set state
	// buy
	profitBuy, profitSale, profitRest := math.MinInt64, math.MinInt64, math.MinInt64
	if buyChances > 0 && status == "cash" {
		profitBuy = maxProfitBackTrace2(prices, day+1, buyChances-1, "stock", profit-prices[day])
	}

	if status == "stock" {
		profitSale = maxProfitBackTrace2(prices, day+1, buyChances, "cash", profit+prices[day])
	}

	profitRest = maxProfitBackTrace2(prices, day+1, buyChances, status, profit)

	return max(max(profitBuy, profitSale), profitRest)
}

func maxProfitBackTrace3(prices []int, day int, buyChances int, status string) int {
	//terminator
	// if buyChances == 0 || day < 0 {
	// 	return 0
	// }

	// if buyChances == 0 {

	// }

	if buyChances == 0 {
		return 0
	}

	if day == 0 {
		if status == "stock" {
			return -prices[0]
		} else {
			return 0
		}
	}
	fmt.Println(day, status)
	// process current level: make a choice and set state
	profitSale, profitBuy, profitRest := math.MinInt64, math.MinInt64, math.MinInt64
	if status == "cash" {
		profitSale = maxProfitBackTrace3(prices, day-1, buyChances, "stock") + prices[day]
	}

	if status == "stock" {
		profitBuy = maxProfitBackTrace3(prices, day-1, buyChances-1, "cash") - prices[day]
	}

	profitRest = maxProfitBackTrace3(prices, day-1, buyChances, status)

	return max(profitSale, max(profitBuy, profitRest))
}

func maxProfit(prices []int) int {
	// return maxProfitBackTrace2(prices, 0, len(prices), "cash", 0)
	return maxProfitBackTrace3(prices, len(prices)-1, 1, "cash")

}

// @lc code=end
