package problem33

// Task: https://leetcode.com/problems/search-in-rotated-sorted-array/

// Solution: https://leetcode.com/problems/search-in-rotated-sorted-array/submissions/1069334276/

// tags:
// #Array
// #Binary Search

// Time: O(logN)
// Space: O(1)
func search(nums []int, target int) int {
	// find pivot. if there isn't, pivot = 0
	// [4,5,6,7,0,1,2] - in this example pivot = 4 (idx by value 0)
	pivot := findPivot(nums)
	// now we have 2 arrays
	// right array: from pivot to array - [0,1,2]
	// left array: from start to pivot - [4,5,6,7]

	// if found in right array, return 'pivot+founded'
	if result := binarySearch(nums[pivot:], target); result != -1 {
		return pivot + result
	}
	// else return result from left array
	return binarySearch(nums[:pivot], target)
}

func findPivot(arr []int) int {
	l, r := 0, len(arr)-1
	for l < r {
		m := l + (r-l)/2
		mVal := arr[m]
		if mVal > arr[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return l
}

func binarySearch(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		m := l + (r-l)/2
		mVal := arr[m]
		if mVal < target {
			l = m + 1
		} else if mVal > target {
			r = m - 1
		} else {
			return m
		}
	}

	return -1
}
