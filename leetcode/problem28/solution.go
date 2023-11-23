package problem28

// Task: https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/

// Solution: https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/submissions/1085617525/

// tags:
// #Array
// #Two Pointers

// Time: O(N*len(needle))
// Space: O(1)
func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
