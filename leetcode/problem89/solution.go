package problem89

// Task: https://leetcode.com/problems/gray-code/

// Solution: https://leetcode.com/problems/gray-code/submissions/1104440187/

// tags:
// #Bit Manipulation

// Time: O(2^n)
// Space: O(1)
func grayCode(n int) []int {
	count := 1 << n

	result := make([]int, 0, count)

	result = append(result, 0, 1)

	k := 1

	for k < n {
		mask := 1 << k
		for i := len(result) - 1; i >= 0; i-- {
			oldD := result[i]
			newD := oldD | mask
			result = append(result, newD)
		}
		k++
	}

	return result
}
