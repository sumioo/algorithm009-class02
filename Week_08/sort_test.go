package week8

import (
	"math/rand"
	"sort"
	"testing"
)

func TestSelectSort(t *testing.T) {
	nums := []int{10, 3, 4, 11, 1, 2, 4, 5}
	selectSort(nums)
	t.Log(nums)
}

func TestBubbleSort(t *testing.T) {
	nums := []int{10, 3, 4, 11, 1, 2, 4, 5}
	bubbleSort(nums)
	t.Log(nums)
}

func TestInsertSort(t *testing.T) {
	nums := []int{10, 3, 4, 11, 1, 2, 4, 5}
	insertSort(nums)
	t.Log(nums)
}

func TestPartition1(t *testing.T) {
	nums := []int{3, 4, 11, 1, 2, 4, 5}
	p := partition1(nums, 0, len(nums)-1)
	t.Log(p, nums)

	nums = []int{3, 4, 5, 15, 25, 45, 55}
	p = partition1(nums, 0, len(nums)-1)
	t.Log(p, nums)
}

func TestPartition(t *testing.T) {
	nums := []int{3, 4, 11, 1, 2, 4, 5}
	p := partition(nums, 0, len(nums)-1)
	t.Log(p, nums)

	nums = []int{3, 4, 5, 15, 25, 45, 55}
	p = partition(nums, 0, len(nums)-1)
	t.Log(p, nums)
}

func TestQuickSort(t *testing.T) {
	nums := []int{3, 4, 11, 1, 2, 4, 5}
	quickSort(nums)
	t.Log(nums)

	nums = []int{1, 2, 3, 4, 5}
	quickSort(nums)
	t.Log(nums)

	nums = []int{0, 0, 1, 2, 3, 5, 6, 1}
	quickSort(nums)
	t.Log(nums)
}

func TestQuickSortP(t *testing.T) {
	nums := generateRandomNums(100000000000)
	quickSortP(nums)
	t.Log(nums)

}

func generateRandomNums(n int) []int {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = rand.Intn(1000)
	}
	return nums
}

func BenchmarkInsertSort(b *testing.B) {
	nums := generateRandomNums(1000)
	for i := 0; i < b.N; i++ {
		insertSort(nums)
	}
}

func BenchmarkSelectSort(b *testing.B) {
	nums := generateRandomNums(1000)
	for i := 0; i < b.N; i++ {
		selectSort(nums)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	nums := generateRandomNums(1000)
	for i := 0; i < b.N; i++ {
		bubbleSort(nums)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	nums := generateRandomNums(10000)
	for i := 0; i < b.N; i++ {
		quickSort(nums)
	}
}

func BenchmarkQuickSortB(b *testing.B) {
	nums := generateRandomNums(100000)
	for i := 0; i < b.N; i++ {
		quickSortB(nums)
	}
}

func BenchmarkQuickSortP(b *testing.B) {
	nums := generateRandomNums(10000000000)
	for i := 0; i < b.N; i++ {
		quickSortP(nums)
	}
}

func TestMerge(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, -1, 2, 3, 4, 5, 7}
	mergeA(nums, 0, 5, 11)
	t.Log(nums)

	mergeB(nums, 0, 5, 11)
	t.Log(nums)
}

func BenchmarkMergeA(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, -1, 2, 3, 4, 5, 7}
	for i := 0; i < b.N; i++ {
		mergeA(nums, 0, 5, 11)
	}
}

func BenchmarkMergeB(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, -1, 2, 3, 4, 5, 7}
	for i := 0; i < b.N; i++ {
		mergeB(nums, 0, 5, 11)
	}
}

func TestMergeSort(t *testing.T) {
	// nums := generateRandomNums(200)
	nums := []int{1, 2, 3, 4, 5, 6, -1, 2, 3, 4, 5, 7}
	mergeSort(nums)
	t.Log(sort.IntsAreSorted(nums))
}

func TestMergeSortB(t *testing.T) {
	// nums := generateRandomNums(200)
	nums := []int{2, 1, 4, 3, 6, 5, -1, 2, 4, 3, 9, 7}
	mergeSortBU(nums)
	t.Log(sort.IntsAreSorted(nums), nums)
}
