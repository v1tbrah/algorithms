package problem191

// Task: https://leetcode.com/problems/number-of-1-bits/

// Solution: https://leetcode.com/problems/number-of-1-bits/submissions/1094620353/

// tags:
// #Bit Manipulation

// Time: O(1)
// Space: O(1)
func hammingWeight(num uint32) int {
	var bits int
	var mask uint32 = 1
	for i := 0; i < 32; i++ {
		if num&mask != 0 {
			bits++
		}
		mask <<= 1
	}
	return bits
}
