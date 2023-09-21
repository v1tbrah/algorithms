package the_k_weakest_rows_in_a_matrix

import "sort"

// Task: https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/

// Solution: https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/submissions/1055590331/

// tags:
// #Array
// #Sorting
// #Matrix

// Idea:
// 1. Make map[count][]idxs
// 2. Make []counts
// 3. Sort counts
// 4. Get first k idxs from map

// Time: O(m*(n+logm))
// Space: O(m)
func kWeakestRows(mat [][]int, k int) []int {
	countToIdxs := make(map[int][]int)
	for i := 0; i < len(mat); i++ {
		var soldiers int
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				soldiers++
			} else {
				break
			}
		}
		countToIdxs[soldiers] = append(countToIdxs[soldiers], i)
	}

	counts := make([]int, 0)
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
