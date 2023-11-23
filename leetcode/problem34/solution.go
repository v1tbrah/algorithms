package problem34

// Task: https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

// Solution: https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/submissions/1105082578/

// tags:
// #Array
// #Binary Search

// Time: O(logN)
// Space: O(1)
func searchRange(nums []int, target int) []int {
	tgIdx := binarySearch(nums, target)
	if tgIdx == -1 {
		return []int{-1, -1}
	}

	sIdx := binarySearchStart(nums[:tgIdx+1], target)

	eInx := binarySearchEnd(nums[tgIdx:], target)

	return []int{sIdx, tgIdx + eInx}
}

func binarySearchEnd(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l+1 < r {
		m := l + (r-l)/2
		if arr[m] > target {
			r = m
		} else {
			l = m
		}
	}

	if r < len(arr) && arr[r] == target {
		return r
	}

	return l
}

func binarySearchStart(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l < r {
		m := l + (r-l)/2
		if arr[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}

	return r
}

func binarySearch(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		m := l + (r-l)/2
		if arr[m] < target {
			l = m + 1
		} else if arr[m] > target {
			r = m - 1
		} else {
			return m
		}
	}

	return -1
}
