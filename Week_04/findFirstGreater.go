package week4

func findFirstGreaterThen(nums []int, x int) int {
	if nums[len(nums)-1] < x {
		return -1
	}
	left, right := 0, len(nums)-1
	y := -1
	for left <= right {
		mid := (left + right) / 2

		if nums[mid] > x {
			y = nums[mid]
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return y
}
