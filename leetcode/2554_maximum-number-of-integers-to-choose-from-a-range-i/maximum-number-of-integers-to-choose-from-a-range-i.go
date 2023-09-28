package maximum_number_of_integers_to_choose_from_a_range_i

// Task: https://leetcode.com/problems/maximum-number-of-integers-to-choose-from-a-range-i/

// Solution: https://leetcode.com/problems/maximum-number-of-integers-to-choose-from-a-range-i/submissions/1061569823/

// tags:
// #Array
// #Hash Table

// Time: O(n)
// Space: O(len(banned))
func maxCount(banned []int, n int, maxSum int) int {
	bannedMap := make(map[int]struct{}, len(banned))
	for _, b := range banned {
		bannedMap[b] = struct{}{}
	}

	var sum int
	var count int
	for i := 1; i <= n; i++ {
		if _, ok := bannedMap[i]; ok {
			continue
		}

		sum += i
		if sum > maxSum {
			return count
		}

		count++
	}

	return count
}
