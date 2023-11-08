package problem14

// Task: https://leetcode.com/problems/longest-common-prefix/

// Solution: https://leetcode.com/problems/longest-common-prefix/submissions/1094678269/

// tags:
// #String

// Time: O(M*N), where M is count of string, N is max length of strings
// Space: O(1)
func longestCommonPrefix(strs []string) string {
	var nextIdx int
	var emptyByte byte
	for {
		var currChar byte
		for _, str := range strs {
			if len(str) <= nextIdx {
				return str[:nextIdx]
			}
			if currChar == emptyByte {
				currChar = str[nextIdx]
			}
			if str[nextIdx] != currChar {
				return str[:nextIdx]
			}
		}
		nextIdx++
	}
}
