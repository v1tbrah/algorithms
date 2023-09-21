package linear_search_and_map

import "sort"

// Task: https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/

// Solution: https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/submissions/1055747546/

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
	newMat := make([][2]int, len(mat))
	for i := 0; i < len(mat); i++ {
		var soldiers int
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				soldiers++
			} else {
				break
			}
		}
		newMat[i] = [2]int{soldiers, i}
	}

	sort.Slice(newMat, func(i, j int) bool {
		return less(newMat[i], newMat[j])
	})

	result := make([]int, 0, k)
	for i := 0; i < k; i++ {
		result = append(result, newMat[i][1])
	}

	return result
}

func less(a, b [2]int) bool {
	if a[0] == b[0] {
		return a[1] < b[1]
	}
	return a[0] < b[0]
}
