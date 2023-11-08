package problem191

// Task: https://leetcode.com/problems/number-of-1-bits/

// Solution: https://leetcode.com/problems/number-of-1-bits/submissions/1094609109/

// tags:
// #Bit Manipulation

// Time: O(1)
// Space: O(1)
func hammingWeight(num uint32) int {
	var count int
	for num > 0 {
		if num%2 == 1 {
			count++
		}
		num /= 2
	}
	return count
}
