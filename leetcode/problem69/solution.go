package problem69

// Task: https://leetcode.com/problems/sqrtx/

// Solution: https://leetcode.com/problems/sqrtx/submissions/1107058665/

// tags:
// #Binary Search
// #Math

// Time: O(LogN)
// Space: O(1)
func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	l, r := 2, x/2
	for l <= r {
		m := l + (r-l)/2
		mVal := m * m

		if mVal < x {
			l = m + 1
		} else if mVal > x {
			r = m - 1
		} else {
			return m
		}
	}

	return l - 1
}
