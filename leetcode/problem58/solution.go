package problem58

// Task: https://leetcode.com/problems/length-of-last-word/

// Solution: https://leetcode.com/problems/length-of-last-word/submissions/1106358834/

// tags:
// #String

// Time: O(N)
// Space: O(1)
func lengthOfLastWord(s string) int {
	var (
		res                       int
		firstNonSpaceAlreadyFound bool
	)

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if firstNonSpaceAlreadyFound {
				return res
			}
		} else {
			firstNonSpaceAlreadyFound = true
			res++
		}
	}

	return res
}
