package sqrtx

func mySqrt(x int) int {
	left, right := 1, x
	for left <= right {
		mid := left + (right-left)/2
		switch {
		case mid*mid > x:
			if (mid-1)*(mid-1) < x {
				return mid - 1
			}
			right = mid - 1
		case mid*mid < x:
			if (mid+1)*(mid+1) > x {
				return mid
			}
			left = mid + 1
		default:
			return mid
		}
	}
	return 0
}
