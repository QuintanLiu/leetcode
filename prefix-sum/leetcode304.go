package prefix_sum

type NumMatrix struct {
	preSum [][]int
}

func Constructor2(matrix [][]int) NumMatrix {
	preSum := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		length := len(matrix[i])
		preSum[i] = make([]int, length)
		if i == 0 {
			for j := 0; j < length; j++ {
				if j == 0 {
					preSum[i][j] = matrix[i][j]
				} else {
					preSum[i][j] = preSum[i][j-1] + matrix[i][j]
				}
			}
		} else {
			for j := 0; j < length; j++ {
				if j == 0 {
					preSum[i][j] = preSum[i-1][j] + matrix[i][j]
				} else {
					preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] - preSum[i-1][j-1] + matrix[i][j]
				}
			}
		}
	}

	return NumMatrix{
		preSum: preSum,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	x, y, z, k := this.preSum[row2][col2], 0, 0, 0
	if row1 == 0 {
		z = 0
		k = 0
	} else {
		z = this.preSum[row1-1][col2]
	}
	if col1 == 0 {
		y = 0
		k = 0
	} else {
		y = this.preSum[row2][col1-1]
	}
	if row1 != 0 && col1 != 0 {
		k = this.preSum[row1-1][col1-1]
	}
	return x - y - z + k
}
