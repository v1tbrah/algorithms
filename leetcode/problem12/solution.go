package problem12

import "strings"

// Task: https://leetcode.com/problems/integer-to-roman/

// Solution: https://leetcode.com/problems/integer-to-roman/submissions/1094599489/

// tags:
// #Hash Table
// #String

// Time: O(1)
// Space: O(1)
func intToRoman(num int) string {
	romanToInt := map[string]int{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}

	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := make([]string, 0)
	for i := 0; i < len(romans) && num > 0; i++ {
		count := num / romanToInt[romans[i]]
		for count > 0 {
			result = append(result, romans[i])
			count--
		}

		num %= romanToInt[romans[i]]
	}

	return strings.Join(result, "")
}
