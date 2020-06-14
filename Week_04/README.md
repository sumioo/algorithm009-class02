# 学习笔记

## 使用二分查找，寻找一个半有序数组 [4, 5, 6, 7, 0, 1, 2] 中间无序的地方

二分查找使用条件有
* 单调性
* 有界
* 索引

二分查找的目的是缩小查找范围，在这里我们要查找数组中间无序的分界点，细心观察可知，假如mid == index(7)，如果 index(mid+1) < index(mid) 或者 mid == index(0) 如果index(mid-1) > index(mid), 那么我们找到了该分界点。

那么依据什么来缩小范围呢？也就是我们依据什么来确定 left = mid + 1 或者 right = mid - 1? 如果nums[0] < nums[mid]那么我们就可以认为这一区间是单调递增的，那么分界点一定是在mid右边,即left = mid + 1，否则就在左边。

这里要注意一点的是，一开始可能只注意到分界点具有的特征是index(mid+1) < index(mid)，但我们应该知道mid也可能落在index(0), 如果不加上这个特征判断，那么当mid == index(0)时，我们会把范围缩小到 [left, index(0)-1]，这样就无法找到分界点了

```go
func search(nums []int, target int) int {
    left, right := 0, len(nums)-1
    for left <= right {
        mid := left + (right - left) / 2
        if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		if nums[mid] < nums[mid-1] {
			return nums[mid]
		}

        if nums[0] < nums[mid] {
            left = mid+1
        }else{
            right = mid - 1
        }
    }
    return -1
}
```