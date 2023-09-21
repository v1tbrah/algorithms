package the_k_weakest_rows_in_a_matrix

// Task: https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/

// Solution: https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/submissions/1055578049/

// tags:
// #Array
// #Sorting
// #Matrix

// Idea:
// 1. Make new matrix, when each row is array[2], where array[0] - count soldiers, array[1] - row idx
// 2. Sort new matrix
// 3. Get idxs from first K in new matrix

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

	for i := 1; i < len(newMat); i++ {
		for j := i; j > 0; j-- {
			if less(newMat[j], newMat[j-1]) {
				newMat[j], newMat[j-1] = newMat[j-1], newMat[j]
			} else {
				break
			}
		}
	}

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
