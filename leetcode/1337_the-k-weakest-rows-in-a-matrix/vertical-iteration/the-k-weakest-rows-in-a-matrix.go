package vertical_iteration

// Task: https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/

// Solution: https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/submissions/1056649826/

// tags:
// #Array
// #Sorting
// #Matrix

// Idea:
// 1. Iterate vertical and add idx if cell==0

// Time: O(m*n)
// Space: O(1)
func kWeakestRows(mat [][]int, k int) []int {
	result := make([]int, 0, k)

	addedIdxs := make(map[int]bool)

	for i := 0; i < len(mat[0]) && len(result) < k; i++ {
		for j := 0; j < len(mat); j++ {
			if mat[j][i] == 0 && !addedIdxs[j] {
				result = append(result, j)
				addedIdxs[j] = true
			}
			if len(result) == k {
				return result
			}
		}
	}

	// if some rows, which have only '1', must be in K weakest rows
	for i := 0; len(result) < k; i++ {
		if mat[i][0] == 1 && !addedIdxs[i] {
			result = append(result, i)
			addedIdxs[i] = true
		}
	}

	return result
}
