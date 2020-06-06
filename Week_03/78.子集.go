/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 *
 * https://leetcode-cn.com/problems/subsets/description/
 *
 * algorithms
 * Medium (77.24%)
 * Likes:    585
 * Dislikes: 0
 * Total Accepted:    93.2K
 * Total Submissions: 120.5K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
 *
 * 说明：解集不能包含重复的子集。
 *
 * 示例:
 *
 * 输入: nums = [1,2,3]
 * 输出:
 * [
 * ⁠[3],
 * [1],
 * [2],
 * [1,2,3],
 * [1,3],
 * [2,3],
 * [1,2],
 * []
 * ]
 *
 */

// @lc code=start

/*
可将回溯看成len(nums)个阶段，每个阶段有两种选择 1 选取当前数字 2 不选取
回溯过程可以看成从左到右扫描数组数字，然后每个数字可以选取或不选取，当index == len(nums)则扫描结束。
*/

func backtrace(nums []int, collection []int, index int, result *[][]int) {
	if len(nums) == index {
		*result = append(*result, append([]int{}, collection...))
		return
	}
	backtrace(nums, collection, index+1, result)
	collection = append(collection, nums[index])
	backtrace(nums, collection, index+1, result)
}

func subsetsA(nums []int) [][]int {
	result := [][]int{}
	collection := []int{}
	backtrace(nums, collection, 0, &result)
	return result
}

func subsets(nums []int) [][]int {
	result := [][]int{}
	result = append(result, []int{})
	for _, num := range nums {
		for _, collection := range result {
			cp := append([]int{}, collection...)
			result = append(result, append(cp, num))
		}
	}
	return result
}

// @lc code=end

