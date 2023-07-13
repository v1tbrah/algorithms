package contains_duplicate

// Task: https://leetcode.com/problems/contains-duplicate/

// Solution: https://leetcode.com/problems/contains-duplicate/submissions/993774943/

// tags:
// #Array
// #Hash Table
// #Sorting

func containsDuplicate(nums []int) bool {
	numSet := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := numSet[num]; ok {
			return true
		}
		numSet[num] = struct{}{}
	}
	return false
}
