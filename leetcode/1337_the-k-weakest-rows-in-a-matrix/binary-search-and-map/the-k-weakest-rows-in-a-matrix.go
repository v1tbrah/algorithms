package binary_search_and_map

import "sort"

// Task: https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/

// Solution: https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/submissions/1055753180/

// tags:
// #Array
// #Sorting
// #Matrix
// #Binary Search

// Idea:
// 1. Make map[count][]idxs, use binarySearch for calculate count on every iteration
// 2. Make []counts
// 3. Sort counts
// 4. Get first k idxs from map

// Time: O(m*logmn)
// Space: O(m)
func kWeakestRows(mat [][]int, k int) []int {
	countToIdxs := make(map[int][]int)
	for i := 0; i < len(mat); i++ {
		count := binarySearch(mat[i])
		countToIdxs[count] = append(countToIdxs[count], i)
	}

	counts := make([]int, 0, len(countToIdxs))
	for c := range countToIdxs {
		counts = append(counts, c)
	}

	sort.Slice(counts, func(i, j int) bool {
		return counts[i] < counts[j]
	})

	result := make([]int, 0, k)
	for _, c := range counts {
		idxs := countToIdxs[c]
		for _, idx := range idxs {
			result = append(result, idx)
			if len(result) == k {
				return result
			}
		}
		if len(result) == k {
			return result
		}
	}

	return result
}

func binarySearch(arr []int) int {
	l, r := 0, len(arr)
	for l < r {
		mid := l + (r-l)/2
		if arr[mid] == 1 {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
