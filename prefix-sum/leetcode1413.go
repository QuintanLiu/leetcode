package prefix_sum

func minStartValue(nums []int) int {
	min := nums[0]
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum < min {
			min = sum
		}
	}
	if min < 0 {
		return 1 - min
	} else {
		return 1
	}
}
