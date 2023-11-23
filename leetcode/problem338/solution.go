package problem338

// Task: https://leetcode.com/problems/counting-bits/

// Solution: https://leetcode.com/problems/counting-bits/submissions/1104912381/

// tags:
// #Bit Manipulation

// Time: O(1)
// Space: O(1)
func countBits(n int) []int {
	bitMasks := make([]int, 32)
	for i := 0; i < len(bitMasks); i++ {
		bitMasks[i] = 1 << i
	}

	result := make([]int, 0, n+1)
	var curr int
	for len(result) < n+1 {
		var bits int
		for i := 0; i < len(bitMasks) && curr >= bitMasks[i]; i++ {
			if bitMasks[i]&curr != 0 {
				bits++
			}
		}
		result = append(result, bits)
		curr++
	}

	return result
}
