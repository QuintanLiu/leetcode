package binary_search

/**
在排序数组中查找元素的第一个和最后一个位置
*/
func searchRange(nums []int, target int) []int {
	left := searchLeft(nums, target)
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	right := searchRight(nums, target)
	return []int{left, right}
}
