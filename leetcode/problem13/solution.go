package problem13

// Task: https://leetcode.com/problems/roman-to-integer/

// Solution: https://leetcode.com/problems/roman-to-integer/submissions/1094672919/

// tags:
// #Hash Table
// #String

// Time: O(1)
// Space: O(1)
func romanToInt(s string) int {
	charToVal := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sum := charToVal[s[0]]
	for i := 1; i < len(s); i++ {
		if charToVal[s[i]] > charToVal[s[i-1]] {
			sum -= 2 * charToVal[s[i-1]]
		}
		sum += charToVal[s[i]]
	}

	return sum
}
