package problem1

// Task: https://leetcode.com/problems/two-sum/

// Solution: https://leetcode.com/problems/two-sum/submissions/1085501208/

// tags:
// #Array
// #Hash Table

// Time: O(N)
// Space: O(N)
func twoSum(nums []int, target int) []int {
	mapNums := make(map[int]int)
	for i, n := range nums {
		if j, ok := mapNums[target-n]; ok && i != j {
			return []int{i, j}
		}
		mapNums[n] = i
	}
	return []int{-1, -1}
}
