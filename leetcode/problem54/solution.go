package problem54

// Task: https://leetcode.com/problems/spiral-matrix/

// Solution: https://leetcode.com/problems/spiral-matrix/submissions/1067936496/

// tags:
// #Array
// #Matrix

// Time: O(M*N)
// Space: O(1)
func spiralOrder(matrix [][]int) []int {
	lenR := len(matrix)
	lenC := len(matrix[0])

	result := make([]int, 0, lenR*lenC)

	rightBorder := lenC - 1
	downBorder := lenR - 1
	leftBorder := 0
	upBorder := 0

	for len(result) != lenR*lenC {
		for i := leftBorder; i <= rightBorder && len(result) != lenR*lenC; i++ {
			result = append(result, matrix[upBorder][i])
		}
		upBorder++

		for i := upBorder; i <= downBorder && len(result) != lenR*lenC; i++ {
			result = append(result, matrix[i][rightBorder])
		}
		rightBorder--

		for i := rightBorder; i >= leftBorder && len(result) != lenR*lenC; i-- {
			result = append(result, matrix[downBorder][i])
		}
		downBorder--

		for i := downBorder; i >= upBorder && len(result) != lenR*lenC; i-- {
			result = append(result, matrix[i][leftBorder])
		}
		leftBorder++
	}

	return result
}
