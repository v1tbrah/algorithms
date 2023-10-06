package task1572

// Task: https://leetcode.com/problems/matrix-diagonal-sum/

// Solution: https://leetcode.com/problems/matrix-diagonal-sum/submissions/1068885749/

// tags:
// #Array
// #Matrix

// Time: O(m)
// Space: O(1)
func diagonalSum(mat [][]int) int {
	var sum int
	for i, j := 0, 0; i < len(mat); i, j = i+1, j+1 {
		pDx, pDy := i, j
		sum += mat[pDx][pDy]
		sDx, sDy := pDx, len(mat)-1-pDy
		if pDx == sDx && pDy == sDy {
			continue
		}
		sum += mat[sDx][sDy]
	}
	return sum
}
