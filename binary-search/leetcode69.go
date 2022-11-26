package binary_search

func mySqrt(x int) int {
	left, right := 0, x
	for left <= right {
		mid := (left + right) >> 1
		m := mid * mid
		if m == x {
			return mid
		} else if m > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}
