package prefix_sum

/**
一维前缀和初始化
*/
func onePrefixSum(arr []int) {
	preSum := make([]int, len(arr))
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		preSum[i] = sum
	}
}

// TwoPrefixSum /**
func TwoPrefixSum(arr [][]int) {
	preSum := make([][]int, len(arr))
	for i := 0; i < len(arr); i++ {
		length := len(arr[i])
		preSum[i] = make([]int, length)
		if i == 0 {
			for j := 0; j < length; j++ {
				if j == 0 {
					preSum[i][j] = arr[i][j]
				} else {
					preSum[i][j] = preSum[i][j-1] + arr[i][j]
				}
			}
		} else {
			for j := 0; j < length; j++ {
				if j == 0 {
					preSum[i][j] = preSum[i-1][j] + arr[i][j]
				} else {
					preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] - preSum[i-1][j-1] + arr[i][j]
				}
			}
		}
	}
}
