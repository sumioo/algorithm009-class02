package week4

import "testing"

func TestFindFirstGreater(t *testing.T) {
	nums := []int{1, 3, 4, 5, 8, 10, 20}
	m := findFirstGreaterThen(nums, 5)
	t.Log(m)
}
