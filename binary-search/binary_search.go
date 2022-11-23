package binary_search

/**
二分查找的一般形式
适用于递增数组
*/
func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

/**
二分查找搜索左边界
注意点：
1、循环跳出的条件
2、目标值被找到的操作
3、left位置的值不一定是target，但是如果target不存在target的值一定介于(right, left]
*/
func searchLeft(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

/**
二分查找搜索右边界
注意点：
1、循环跳出的条件
2、目标值被找到的操作
3、right位置的值不一定是target，但是如果target不存在target的值一定介于[right, left)
*/
func searchRight(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}
