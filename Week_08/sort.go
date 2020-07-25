package week8

func selectSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}

func bubbleSort(nums []int) {
	for i := len(nums) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

/*
从前到后逐步构建有序序列
假设前n位是有序的，那么n+1位和前n位比较，找出小于他的那个位置，在那个位置后插入
起始时假设前1个是有序的
如果序列是大部分有序的，插入排序有较好的性能，选择排序和冒泡排序不具有这个特点
*/
func insertSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		j, curr := i, nums[i]
		for ; j > 0 && nums[j-1] > curr; j-- {
			nums[j] = nums[j-1]
		}
		nums[j] = curr
	}
}

/*
i指针从左到右扫描，直到找到一个大于piviot的数字，j指针从右到左扫描直到找到小于piviot的数字，交换两个指针指向的数字，继续，直到left >= right

1. 数组有序递增
2. 数组有序递减
3. 普通情况
*/
func partition1(nums []int, left int, right int) int {
	piviot := left
	left++
	for {
		for left < right && nums[left] <= nums[piviot] {
			left++
		}

		for right >= left && nums[right] > nums[piviot] { //>=表明right最终可能的取值可为left-1
			right--
		}

		if left >= right {
			break
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	nums[piviot], nums[right] = nums[right], nums[piviot]
	return right
}

/*
主指针i负责遍历数组，指针j指向大于pivot的元素
*/
func partition(nums []int, left int, right int) int {
	pivot := right
	counter := left
	for i := left; i < right; i++ {
		if nums[i] <= nums[pivot] {
			// if counter != i {
			nums[counter], nums[i] = nums[i], nums[counter]
			// }
			counter++
		}
	}
	nums[counter], nums[pivot] = nums[pivot], nums[counter]
	return counter
}

func _quickSort(nums []int, left int, right int) {
	if left >= right {
		return
	}
	piviot := partition(nums, left, right)
	_quickSort(nums, left, piviot-1)
	_quickSort(nums, piviot+1, right)
}

func _quickSortB(nums []int, left int, right int) {
	if left >= right {
		return
	}
	piviot := partition1(nums, left, right)
	_quickSortB(nums, left, piviot-1)
	_quickSortB(nums, piviot+1, right)
}

func quickSort(nums []int) {
	_quickSort(nums, 0, len(nums)-1)
}

func quickSortB(nums []int) {
	_quickSortB(nums, 0, len(nums)-1)
}

func quickSortP(nums []int) {
	left, right := 0, len(nums)-1
	piviot := partition1(nums, left, right)
	go _quickSortB(nums, left, piviot-1)
	go _quickSortB(nums, piviot+1, right)
}

func mergeA(nums []int, left int, mid int, right int) {
	aux := make([]int, right-left+1)
	for i := left; i <= right; i++ {
		aux[i] = nums[i]
	}

	i, j := left, mid+1
	for k := left; k <= right; k++ {
		if i > mid {
			nums[k] = aux[j]
			j++
		} else if j > right {
			nums[k] = aux[i]
			i++
		} else if aux[i] < aux[j] {
			nums[k] = aux[i]
			i++
		} else {
			nums[k] = aux[j]
			j++
		}
	}
}

func mergeB(nums []int, left int, mid int, right int) {
	aux := make([]int, right-left+1)
	i, j, k := left, mid+1, 0
	for i <= mid && j <= right {
		if nums[i] < nums[j] {
			aux[k] = nums[i]
			i++
		} else {
			aux[k] = nums[j]
			j++
		}
		k++
	}

	for i <= mid {
		aux[k] = nums[i]
		i++
		k++
	}
	for j <= right {
		aux[k] = nums[j]
		j++
		k++
	}

	for i := range aux {
		nums[left+i] = aux[i]
	}
}

func _mergeSort(nums []int, left int, right int) {
	if left >= right {
		return
	}
	mid := left + (right-left)/2
	_mergeSort(nums, left, mid)
	_mergeSort(nums, mid+1, right)
	mergeB(nums, left, mid, right)
}

func mergeSort(nums []int) {
	_mergeSort(nums, 0, len(nums)-1)
}

//自底向上的归并
func mergeSortBU(nums []int) {
	min := func(x int, y int) int {
		if x > y {
			return y
		}
		return x
	}

	numsLen := len(nums)
	for sz := 1; sz < numsLen; sz += sz {
		for left := 0; left < numsLen; left += sz + sz {
			mergeB(nums, left, left+sz-1, min(left+sz+sz-1, numsLen-1))
		}
	}
}
