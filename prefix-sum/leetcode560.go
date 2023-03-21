package prefix_sum

func subarraySum(nums []int, k int) int {
	res := 0
	pre := 0
	m := map[int]int{}
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		p, ok := m[pre-k]
		if ok {
			res += p
		}
		m[pre] += 1
	}
	return res
}
