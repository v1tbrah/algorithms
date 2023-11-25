package problem56

import "sort"

// Task: https://leetcode.com/problems/merge-intervals/

// Solution: https://leetcode.com/problems/merge-intervals/submissions/1106356855/

// tags:
// #Array
// #Sorting

// Time: O(N*LogN)
// Space: O(N)
func merge(intervals [][]int) [][]int {
	ends, starts := make([]int, 0, len(intervals)), make([]int, 0, len(intervals))
	for _, se := range intervals {
		starts = append(starts, se[0])
		ends = append(ends, se[1])
	}

	sort.Slice(starts, func(i, j int) bool { return starts[i] < starts[j] })
	sort.Slice(ends, func(i, j int) bool { return ends[i] < ends[j] })

	result := make([][]int, 0)

	var leftIdx int
	var currMaxEnd int
	for i := 0; i < len(starts)-1; i++ {
		currMaxEnd = max(currMaxEnd, ends[i])
		if currMaxEnd < starts[i+1] {
			result = append(result, []int{starts[leftIdx], currMaxEnd})
			leftIdx = i + 1
		}
	}

	currMaxEnd = max(currMaxEnd, ends[len(ends)-1])
	result = append(result, []int{starts[leftIdx], currMaxEnd})

	return result
}
