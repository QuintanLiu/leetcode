package prefix_sum

func minSubArrayLen(target int, nums []int) int {
	i, j := 0, 0
	sum := nums[0]
	res := 0
	for i <= j && j < len(nums) {
		if sum < target {
			j++
			if j != len(nums) {
				sum += nums[j]
			}
		} else if sum >= target {
			length := j - i + 1
			if res == 0 || res > length {
				res = length
				if res == 1 {
					break
				}
			}
			sum -= nums[i]
			i++
		}
	}
	return res
}

/**
使用二分查找解决这个问题
*/
func minSubArrayLen2(target int, nums []int) int {
	var arr []int
	sum := 0
	res := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		arr = append(arr, sum)
	}
	if sum < target {
		return 0
	}
	for i := 0; i < len(arr); i++ {
		left, right := i, len(arr)-1
		for left <= right {
			mid := (left + right) >> 1
			var sub int
			if i == 0 {
				sub = 0
			} else {
				sub = arr[i-1]
			}
			if arr[mid]-sub >= target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		if left == len(arr) {
			break
		}
		if res == 0 || res > left-i+1 {
			res = left - i + 1
		}
	}
	return res
}
