package problem27

// Task: https://leetcode.com/problems/remove-element/

// Solution: https://leetcode.com/problems/remove-element/submissions/1085615173/

// tags:
// #Array
// #Two Pointers

// Time: O(N)
// Space: O(1)
func removeElement(nums []int, val int) int {
	insertPosition := 0
	for _, n := range nums {
		if n != val {
			nums[insertPosition] = n
			insertPosition++
		}
	}
	return insertPosition
}
