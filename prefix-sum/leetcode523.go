package prefix_sum

func checkSubarraySum(nums []int, k int) bool {
	preSumMap := make(map[int]int, len(nums))
	preSumArr := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i != 0 {
			preSumArr[i] = nums[i] + preSumArr[i-1]
			if index, has := preSumMap[preSumArr[i]%k]; has {
				if i-index > 2 {
					return true
				}
			} else {
				preSumMap[preSumArr[i]%k] = i
			}
		} else {
			preSumArr[i] = nums[i]
			preSumMap[preSumArr[i]%k] = i
		}
	}

	return false
}
