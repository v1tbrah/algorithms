package longest_substring_with_at_most_k_distinct_characters

// Task: https://leetcode.com/problems/longest-substring-with-at-most-k-distinct-characters/

// Solution: https://leetcode.com/problems/longest-substring-with-at-most-k-distinct-characters/submissions/1067123107/

// tags:
// #Hash Table
// #String
// #Sliding Window

// Time: O(n*k)
// Space: O(k)
func lengthOfLongestSubstringKDistinct(s string, k int) int {
	if k == 0 {
		return 0
	}

	// s = "eceba", k = 2
	// firstCharIdx = 3
	// currLen = 2
	// maxLen = 3
	// mapCurrChars[b: 1, a: 1]
	firstCharIdx := 0
	currLen := 1
	maxLen := 1

	mapCurrChars := make(map[byte]int, k+1)
	mapCurrChars[s[0]]++

	for i := 1; i < len(s); i++ {
		mapCurrChars[s[i]]++
		if len(mapCurrChars) > k {
			if currLen > maxLen {
				maxLen = currLen
			}
			delete(mapCurrChars, s[i])
			for j := firstCharIdx; len(mapCurrChars) == k; j++ {
				mapCurrChars[s[j]]--
				if mapCurrChars[s[j]] == 0 {
					delete(mapCurrChars, s[j])
				}
				firstCharIdx++
			}
			currLen = i - firstCharIdx + 1
			mapCurrChars[s[i]]++
			continue
		}
		currLen++
	}

	if currLen > maxLen {
		maxLen = currLen
	}

	return maxLen
}
