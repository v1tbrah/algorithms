package problem136

// Task: https://leetcode.com/problems/single-number/

// Solution: https://leetcode.com/problems/single-number/submissions/1105076843/

// tags:
// #Array
// #Bit Manipulation

// Time: O(N)
// Space: O(1)
func singleNumber(nums []int) int {
	var res int

	for _, n := range nums {
		res ^= n
	}

	return res
}
