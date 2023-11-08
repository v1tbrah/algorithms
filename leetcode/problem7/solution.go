package problem7

import "math"

// Task: https://leetcode.com/problems/reverse-integer/

// Solution: https://leetcode.com/problems/reverse-integer/submissions/1092295117/

// tags:
// #Math

// Time: O(LgX)
// Space: O(LgX)
func reverse(x int) int {
	digits := make([]int, 0)
	for abs(x) > 9 {
		digits = append(digits, x%10)
		x /= 10
	}
	digits = append(digits, x)

	muller := 1

	var total int
	for i := len(digits) - 1; i >= 0; i-- {
		total += digits[i] * muller
		if total > math.MaxInt32 || total < math.MinInt32 {
			return 0
		}

		muller *= 10
	}

	return total
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
