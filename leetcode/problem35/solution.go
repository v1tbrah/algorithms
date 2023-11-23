package problem35

// Task: https://leetcode.com/problems/search-insert-position/

// Solution: https://leetcode.com/problems/search-insert-position/submissions/1105083016/

// tags:
// #Array
// #Binary Search

// Time: O(logN)
// Space: O(1)
func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			return m
		}
	}

	return r + 1
}
