package problem73

// Task: https://leetcode.com/problems/set-matrix-zeroes/

// Solution: https://leetcode.com/problems/set-matrix-zeroes/submissions/1068586656/

// tags:
// #Matrix
// #Hash Table

// Time: O(M*N)
// Space: O(M+N)
func setZeroes(m [][]int) {
	rows := make(map[int]struct{})
	cols := make(map[int]struct{})

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == 0 {
				rows[i] = struct{}{}
				cols[j] = struct{}{}
			}
		}
	}

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if _, ok := rows[i]; ok {
				m[i][j] = 0
			}
			if _, ok := cols[j]; ok {
				m[i][j] = 0
			}
		}
	}
}
