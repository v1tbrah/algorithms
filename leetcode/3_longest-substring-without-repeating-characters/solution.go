package problem3

// Task: https://leetcode.com/problems/longest-substring-without-repeating-characters/

// Solution: https://leetcode.com/problems/longest-substring-without-repeating-characters/submissions/1091582080/

// tags:
// #Hash Table
// #String
// #Sliding Window

// Time: O(N)
// Space: O(N)
func lengthOfLongestSubstring(s string) int {
	letterToIdx := make(map[byte]int)

	var left, maxLen int
	for i := 0; i < len(s); i++ {
		if j, ok := letterToIdx[s[i]]; ok && j >= left {
			left = j + 1
		}

		letterToIdx[s[i]] = i

		maxLen = max(maxLen, i-left+1)
	}

	return maxLen
}
