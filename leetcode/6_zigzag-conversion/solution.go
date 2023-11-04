package problem6

// Task: https://leetcode.com/problems/zigzag-conversion/

// Solution: https://leetcode.com/problems/zigzag-conversion/submissions/1091584493/

// tags:
// #String
// #Matrix

// Time: O(numRows*numCols)
// Space: O(numRows*numCols)
func convert(s string, numRows int) string {
	matrix := make([][]byte, numRows)

	var currStrIdx int
	for currStrIdx < len(s) {
		for i := 0; i < numRows && currStrIdx < len(s); i++ {
			matrix[i] = append(matrix[i], s[currStrIdx])
			currStrIdx++
		}
		for i := numRows - 1 - 1; i > 0 && currStrIdx < len(s); i-- {
			matrix[i] = append(matrix[i], s[currStrIdx])
			currStrIdx++
		}
	}

	result := make([]byte, 0)
	for _, str := range matrix {
		result = append(result, str...)
	}

	return string(result)
}
