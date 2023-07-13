package intersection_of_two_arrays

// Task: https://leetcode.com/problems/intersection-of-two-arrays/

// Solution: https://leetcode.com/problems/intersection-of-two-arrays/submissions/993807505/

// tags:
// #Array
// #Hash Table
// #Two pointers
// #Binary Search
// #Sorting

func intersection(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	resultMap := make(map[int]struct{}, 0)

	nums1Map := make(map[int]struct{}, len(nums1))
	for _, num := range nums1 {
		nums1Map[num] = struct{}{}
	}

	for _, num := range nums2 {
		if _, ok1 := nums1Map[num]; ok1 {
			if _, ok2 := resultMap[num]; !ok2 {
				result = append(result, num)
				resultMap[num] = struct{}{}
			}
		}
	}

	return result
}
