package problem66

// Task: https://leetcode.com/problems/plus-one/

// Solution: https://leetcode.com/problems/plus-one/submissions/824662824/

// tags:
// #Array
// #Math

// Time: O(N)
// Space: O(1)
func plusOne(digits []int) []int {
	needInc := true
	for j := len(digits) - 1; j >= 0 && needInc; j-- {
		if digits[j] == 9 {
			digits[j] = 0
			needInc = true
		} else {
			digits[j] += 1
			needInc = false
		}
	}
	if needInc {
		digits = append([]int{1}, digits...)
	}
	return digits
}
