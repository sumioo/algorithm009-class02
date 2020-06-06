/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 *
 * https://leetcode-cn.com/problems/permutations/description/
 *
 * algorithms
 * Medium (76.04%)
 * Likes:    728
 * Dislikes: 0
 * Total Accepted:    134.9K
 * Total Submissions: 177.2K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个 没有重复 数字的序列，返回其所有可能的全排列。
 *
 * 示例:
 *
 * 输入: [1,2,3]
 * 输出:
 * [
 * ⁠ [1,2,3],
 * ⁠ [1,3,2],
 * ⁠ [2,1,3],
 * ⁠ [2,3,1],
 * ⁠ [3,1,2],
 * ⁠ [3,2,1]
 * ]
 *
 */

/*
我们可以将排列分为n个阶段，每一个阶段可以选取没有被使用过的数字，所以这就启发我们需要一种数据结构来记录
数字的使用情况。
在进入下一阶段时我们先把使用了的数字标记为已使用，待到回溯归来时再将数字标记为未使用，这样该数字就可以
被另一条路径使用了。

*/

// @lc code=start

func backtrace(nums []int, collection []int, used map[int]bool, result *[][]int) {
	if len(collection) == len(nums) {
		*result = append(*result, append([]int{}, collection...))
		return
	}

	for _, num := range nums {
		if used[num] {
			continue
		}
		//设置状态
		used[num] = true
		collection = append(collection, num)
		//进入下一阶段
		backtrace(nums, collection, used, result)
		//撤销状态
		used[num] = false
		collection = collection[:len(collection)-1]
	}
}
func permute(nums []int) [][]int {
	result := [][]int{}
	used := map[int]bool{}
	for _, num := range nums {
		used[num] = false
	}
	backtrace(nums, []int{}, used, &result)
	return result
}

// @lc code=end

