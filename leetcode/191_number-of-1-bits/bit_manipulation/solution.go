package problem191

// Task: https://leetcode.com/problems/number-of-1-bits/

// Solution: https://leetcode.com/problems/number-of-1-bits/submissions/1094618378/

// tags:
// #Bit Manipulation

// Time: O(1)
// Space: O(1)
func hammingWeight(num uint32) int {
	var bits int
	for num > 0 {
		bits++
		num &= num - 1
	}
	return bits
}
