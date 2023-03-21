package prefix_sum

func numSubarraysWithSum(nums []int, goal int) int {
	pre, res := 0, 0
	m := map[int]int{}
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		_, ok := m[pre-goal]
		if ok {
			res += m[pre-goal]
		}
		m[pre] += 1
	}
	return res
}
