package problem201

// Task: https://leetcode.com/problems/bitwise-and-of-numbers-range/

// Solution: https://leetcode.com/problems/bitwise-and-of-numbers-range/submissions/1104878982/

// tags:
// #Bit Manipulation

// Time: O(1)
// Space: O(1)
func rangeBitwiseAnd(left int, right int) int {
	shifts := 0
	for left != right {
		left >>= 1
		right >>= 1
		shifts++
	}
	return left << shifts
}
