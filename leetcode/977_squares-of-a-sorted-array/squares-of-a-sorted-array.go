package squares_of_a_sorted_array

// Task: https://leetcode.com/problems/squares-of-a-sorted-array/

// Solution: https://leetcode.com/problems/squares-of-a-sorted-array/submissions/1067076116/

// tags:
// #Array
// #Two Pointers

// Time: O(n)
// Space: O(1)
func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))

	l, r, k := 0, len(nums)-1, len(nums)-1
	for k >= 0 {
		sqrtL, sqrtR := sqrt(nums[l]), sqrt(nums[r])
		if sqrtL > sqrtR {
			result[k] = sqrtL
			l++
		} else {
			result[k] = sqrtR
			r--
		}
		k--
	}

	return result
}

func sqrt(a int) int {
	return a * a
}
