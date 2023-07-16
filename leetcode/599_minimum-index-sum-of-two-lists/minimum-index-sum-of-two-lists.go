package minimum_index_sum_of_two_lists

import "math"

// Task: https://leetcode.com/problems/minimum-index-sum-of-two-lists/

// Solution: https://leetcode.com/problems/minimum-index-sum-of-two-lists/submissions/995915176/

// tags:
// #Array
// #Hash Table
// #String

func findRestaurant(list1 []string, list2 []string) []string {
	list1Map := make(map[string]int, len(list1))
	for i, str := range list1 {
		list1Map[str] = i
	}

	minSum := math.MaxInt
	strToSum := make(map[string]int)
	for i, str := range list2 {
		if j, ok := list1Map[str]; ok {
			sum := i + j
			strToSum[str] = sum
			if sum < minSum {
				minSum = sum
			}
		}
	}

	result := make([]string, 0)
	for str, sum := range strToSum {
		if sum == minSum {
			result = append(result, str)
		}
	}

	return result
}
