package twopointers

import "sort"

// Task: https://leetcode.com/problems/meeting-rooms-ii/

// Solution: https://leetcode.com/problems/meeting-rooms-ii/submissions/1058892263/

// tags:
// #Array
// #Two Pointers
// #Sorting

// Time: O(n*logn)
// Space: O(n)
func minMeetingRooms(intervals [][]int) int {
	starts, ends := make([]int, len(intervals)), make([]int, len(intervals))
	for i, interval := range intervals {
		starts[i] = interval[0]
		ends[i] = interval[1]
	}

	sort.Slice(starts, func(i, j int) bool {
		return starts[i] < starts[j]
	})
	sort.Slice(ends, func(i, j int) bool {
		return ends[i] < ends[j]
	})

	var result int
	for startIdx, endIdx := 0, 0; startIdx < len(starts); startIdx++ {
		if starts[startIdx] < ends[endIdx] {
			result++
		} else {
			endIdx++
		}
	}

	return result
}
