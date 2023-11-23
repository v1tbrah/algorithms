package problem26

// Task: https://leetcode.com/problems/remove-duplicates-from-sorted-array/

// Solution: https://leetcode.com/problems/remove-duplicates-from-sorted-array/submissions/1105078838/

// tags:
// #Array
// #Two Pointers

// Time: O(N)
// Space: O(1)
func removeDuplicates(nums []int) int {
	lastNonDuplIdx := 0

	for i := 1; i < len(nums); i++ {
		lastNonDupl := nums[lastNonDuplIdx]
		curr := nums[i]
		isDupl := lastNonDupl == curr
		if !isDupl {
			lastNonDuplIdx++
			nums[lastNonDuplIdx] = curr
		}
	}

	return lastNonDuplIdx + 1
}
