package find_the_prefix_common_array_of_two_arrays

// Task: https://leetcode.com/problems/find-the-prefix-common-array-of-two-arrays/

// Solution: https://leetcode.com/problems/find-the-prefix-common-array-of-two-arrays/submissions/

// tags:
// #Array
// #Hash Table

// Time: O(n)
// Space: O(n)
func findThePrefixCommonArray(A []int, B []int) []int {
	result := make([]int, 0, len(A))

	mapA, mapB := make(map[int]struct{}), make(map[int]struct{})

	var prev int
	for i := range A {
		mapA[A[i]] = struct{}{}
		mapB[B[i]] = struct{}{}

		if A[i] == B[i] {
			prev++
		} else {
			if _, ok := mapA[B[i]]; ok {
				prev++
			}
			if _, ok := mapB[A[i]]; ok {
				prev++
			}
		}

		result = append(result, prev)
	}

	return result
}
