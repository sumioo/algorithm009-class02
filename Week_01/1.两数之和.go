/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for index, num := range nums {
		r := target - num
		rIndex, ok := m[r]
		if ok {
			return []int{rIndex, index}
		} else {
			m[num] = index
		}
	}
	return nil
}

// @lc code=end

