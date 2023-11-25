package problem48

// Task: https://leetcode.com/problems/rotate-image/

// Solution: https://leetcode.com/problems/rotate-image/submissions/1106354151/

// tags:
// #Array
// #Math
// #Matrix

// Time: O(M*N)
// Space: O(M*N)
func rotate(matrix [][]int) {
	transpose(matrix)
	reverse(matrix)
}

func transpose(m [][]int) {
	for i := 0; i < len(m); i++ {
		for j := i; j < len(m[i]); j++ {
			m[i][j], m[j][i] = m[j][i], m[i][j]
		}
	}
}

func reverse(m [][]int) {
	for i := 0; i < len(m); i++ {
		for l, r := 0, len(m[i])-1; l < r; l, r = l+1, r-1 {
			m[i][l], m[i][r] = m[i][r], m[i][l]
		}
	}
}
