package task1672

// Task: https://leetcode.com/problems/richest-customer-wealth/

// Solution: https://leetcode.com/problems/richest-customer-wealth/submissions/1068819267/

// tags:
// #Array
// #Matrix

// Time: O(m*n)
// Space: O(1)
func maximumWealth(accounts [][]int) int {
	var max int
	for i := 0; i < len(accounts); i++ {
		var sum int
		for j := 0; j < len(accounts[i]); j++ {
			sum += accounts[i][j]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
