package sqrtx

import "math"

func mySqrt(x int) int {
	left := 0
	right := int(math.Pow(2.0, 31.0) - 1)
	for left <= right {
		mid := left + (right-left)/2
		square := mid * mid
		if square < x {
			left = mid + 1
		} else if square > x {
			prev := mid - 1
			if prev*prev < x {
				return prev
			}
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
