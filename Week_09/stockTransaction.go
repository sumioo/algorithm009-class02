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

// @lc code=start

package week9

import (
	"fmt"
	"math"
)

func max122(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
没有重复子问题,只能通过剪枝优化，不过好像也没什么可以减枝的
这里是通过暴力搜索，好像有点类似图的搜索的意思，只是不断的枚举不同的选择
*/
func maxProfitBackTraceDFS(prices []int, day int, buyChances int, status string, profit int) int {
	//terminator
	if (buyChances == 0 && status == "cash") || day >= len(prices) {
		return profit
	}

	// process current level: make a choice and set state
	// buy
	profitBuy, profitSale, profitRest := math.MinInt64, math.MinInt64, math.MinInt64
	if buyChances > 0 && status == "cash" {
		profitBuy = maxProfitBackTraceDFS(prices, day+1, buyChances-1, "stock", profit-prices[day])
	}

	if status == "stock" {
		profitSale = maxProfitBackTraceDFS(prices, day+1, buyChances, "cash", profit+prices[day])
	}

	profitRest = maxProfitBackTraceDFS(prices, day+1, buyChances, status, profit)

	return max122(max122(profitBuy, profitSale), profitRest)
}

/*
思路是这样的，假设当前处于最后一天，那么这天的利润肯定依赖于前n天的利润，也就意味着有子问题。因为股票只有卖出才能获得利润，所以
最后一天肯定是持有现金的状态（这里定义两个状态一个是持有股票状态和只有现金状态）。如果当前阶段的状态要求是只有现金，那么我们有两种选择
一个是卖出股票，一个是休息，什么都不做。如果选择是卖出股票，那么我们的子问题必须是持有股票状态下的子问题，如果是选择休息，那么我们的子问题必须是持有现金状态下的子问题
*/

const CASH = 0
const STOCK = 1

type DayStatus [3]int

func maxProfitBackTrace122(prices []int, day int, buyChances int, status int, memo map[DayStatus]int) int {
	if buyChances == 0 {
		return 0
	}

	if day == 0 {
		if status == STOCK {
			return -prices[0]
		}
		return 0
	}

	dayStatus := DayStatus{day, status, buyChances}
	if v, ok := memo[dayStatus]; ok {
		return v
	}
	fmt.Println(day, status)
	// process current level: make a choice and set state
	profitSale, profitBuy, profitRest := math.MinInt64, math.MinInt64, math.MinInt64
	if status == CASH {
		profitSale = maxProfitBackTrace122(prices, day-1, buyChances, STOCK, memo) + prices[day]
	}

	if status == STOCK {
		profitBuy = maxProfitBackTrace122(prices, day-1, buyChances-1, CASH, memo) - prices[day]
	}

	profitRest = maxProfitBackTrace122(prices, day-1, buyChances, status, memo)
	maxProfit := max122(profitSale, max122(profitBuy, profitRest))
	memo[dayStatus] = maxProfit
	return maxProfit
}

func maxProfit122(prices []int) int {
	// return maxProfitBackTrace2(prices, 0, len(prices), "cash", 0)
	memo := map[DayStatus]int{}
	return maxProfitBackTrace122(prices, len(prices)-1, 2, CASH, memo)
}
