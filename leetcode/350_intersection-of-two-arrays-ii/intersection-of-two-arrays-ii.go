package intersection_of_two_arrays_ii

// Task: https://leetcode.com/problems/intersection-of-two-arrays-ii/

// Solution: https://leetcode.com/problems/intersection-of-two-arrays-ii/submissions/996113245/

// tags:
// #Array
// #Hash Table
// #Two pointers
// #Binary Search
// #Sorting

func intersect(nums1 []int, nums2 []int) []int {
	nums1Map := make(map[int]int)
	for _, num := range nums1 {
		nums1Map[num]++
	}

	result := make([]int, 0)
	for _, num := range nums2 {
		if count := nums1Map[num]; count > 0 {
			result = append(result, num)
			nums1Map[num]--
		}
	}

	return result
}
