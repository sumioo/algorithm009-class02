/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 *
 * https://leetcode-cn.com/problems/majority-element/description/
 *
 * algorithms
 * Easy (63.26%)
 * Likes:    617
 * Dislikes: 0
 * Total Accepted:    172.8K
 * Total Submissions: 272.5K
 * Testcase Example:  '[3,2,3]'
 *
 * 给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
 *
 * 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
 *
 *
 *
 * 示例 1:
 *
 * 输入: [3,2,3]
 * 输出: 3
 *
 * 示例 2:
 *
 * 输入: [2,2,1,1,1,2,2]
 * 输出: 2
 *
 *
 */

// @lc code=start

//Boyer-Moore 投票算法
//维护一个major和count变量。如果count == 0 则major为当前数字，如果count不为0，如果当前数字等于major择count++,如果不等于则count--，继续扫描下一个数字。
func majorityElement(nums []int) int {
	major := 0
	count := 0

	for _, num := range nums {
		if count == 0 {
			major = num
		}
		if num != major {
			count--
		} else {
			count++
		}
	}
	return major
}

// @lc code=end

