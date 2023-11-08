package problem9

// Task: https://leetcode.com/problems/palindrome-number/

// Solution: https://leetcode.com/problems/palindrome-number/submissions/1093860426/

// tags:
// #Math

// Time: O(LgX)
// Space: O(LgX)
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	digits := make([]int, 0)
	for x > 0 {
		digits = append(digits, x%10)
		x /= 10
	}

	for l, r := 0, len(digits)-1; l < r; l, r = l+1, r-1 {
		if digits[l] != digits[r] {
			return false
		}
	}

	return true
}
