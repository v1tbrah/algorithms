package summary_ranges

import "strconv"

// Task: https://leetcode.com/problems/summary-ranges/

// Solution: https://leetcode.com/problems/summary-ranges/submissions/1066189099/

// tags:
// #Array

// Time: O(n)
// Space: O(1)
func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}

	result := make([]string, 0)

	start, end := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		curr := nums[i]
		if curr-end > 1 {
			result = append(result, str(start, end))
			start, end = curr, curr
			continue
		}

		end = curr
	}

	result = append(result, str(start, end))

	return result
}

func str(a, b int) string {
	if a == b {
		return strconv.Itoa(a)
	}
	return strconv.Itoa(a) + "->" + strconv.Itoa(b)
}
