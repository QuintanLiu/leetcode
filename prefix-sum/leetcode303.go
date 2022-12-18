package prefix_sum

type NumArray struct {
	preSum []int
}

func Constructor(nums []int) NumArray {
	preArr := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			preArr[i] = nums[i]
		} else {
			preArr[i] = preArr[i-1] + nums[i]
		}
	}
	return NumArray{
		preSum: preArr,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	leftSum := 0
	if left == 0 {
		leftSum = 0
	} else {
		leftSum = this.preSum[left-1]
	}
	return this.preSum[right] - leftSum
}
