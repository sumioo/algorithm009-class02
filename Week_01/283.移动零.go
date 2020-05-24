/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	pos := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[pos] = nums[i]
			pos++
		}
	}

	for ; pos < len(nums); pos++ {
		nums[pos] = 0
	}
}

// @lc code=end

// @lc code=start
func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		if j < i {
			nums[j], nums[i] = nums[i], nums[j]
		}
		j++
	}
}

// @lc code=end