package task1351

// Task: https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix/

// Solution: https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix/submissions/1068889784/

// tags:
// #Array
// #Matrix

// Time: O(m*n)
// Space: O(1)
func countNegatives(grid [][]int) int {
	var count int
	for i := 0; i < len(grid); i++ {
		for j := len(grid[i]) - 1; j >= 0; j-- {
			if grid[i][j] >= 0 {
				break
			}
			count++
		}
	}
	return count
}
