package problem371

// Task: https://leetcode.com/problems/sum-of-two-integers/

// Solution: https://leetcode.com/problems/sum-of-two-integers/submissions/1096442283/

// tags:
// #Bit Manipulation

// Time: O(1)
// Space: O(1)
func getSum(a int, b int) int {
	for b != 0 {
		answer := a ^ b
		carry := (a & b) << 1
		a = answer
		b = carry
	}

	return a
}
