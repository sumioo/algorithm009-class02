/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 *
 * https://leetcode-cn.com/problems/permutations-ii/description/
 *
 * algorithms
 * Medium (58.65%)
 * Likes:    305
 * Dislikes: 0
 * Total Accepted:    61.9K
 * Total Submissions: 105.1K
 * Testcase Example:  '[1,1,2]'
 *
 * 给定一个可包含重复数字的序列，返回所有不重复的全排列。
 *
 * 示例:
 *
 * 输入: [1,1,2]
 * 输出:
 * [
 * ⁠ [1,1,2],
 * ⁠ [1,2,1],
 * ⁠ [2,1,1]
 * ]
 *
 */

// @lc code=start

func bracktrace(uniqueNums []int, leftCount map[int]int, collection []int, n int, result *[][]int) {
	if len(collection) == n {
		*result = append(*result, append([]int{}, collection...))
		return
	}

	for _, num := range uniqueNums {
		if leftCount[num] == 0 {
			continue
		}
		leftCount[num]--
		collection = append(collection, num)
		bracktrace(uniqueNums, leftCount, collection, n, result)
		leftCount[num]++
		collection = collection[:len(collection)-1]
	}
}

func permuteUnique(nums []int) [][]int {
	result := [][]int{}
	leftCount := map[int]int{}
	for _, num := range nums {
		leftCount[num]++
	}
	uniqueNums := []int{}
	for num := range leftCount {
		uniqueNums = append(uniqueNums, num)
	}
	bracktrace(uniqueNums, leftCount, []int{}, len(nums), &result)
	return result
}

// @lc code=end

