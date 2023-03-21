package prefix_sum

func findMaxLength(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i] = -1
		}
	}
	m := map[int]int{}
	pre, res := 0, 0
	m[0] = -1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		_, ok := m[pre]
		if ok {
			index := m[pre]
			length := i - index
			if length > res {
				res = length
			}
		} else {
			m[pre] = i
		}
	}
	return res
}
