/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSumA(nums []int, target int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSumB(nums []int, target int) []int {
	numIndex := map[int]int{}

	for index, num := range nums {
		numIndex[num] = index
	}

	for index, num := range nums {
		if i, ok := numIndex[target-num]; ok {
			return []int{index, i}
		}
	}
	return []int{}
}

func twoSum(nums []int, target int) []int {
	numIndex := map[int]int{}
	for index, num := range nums {
		if i, ok := numIndex[target-num]; ok {
			return []int{i, index}
		}
		numIndex[num] = index
	}
	return []int{}
}

// @lc code=end

