package longest_substring_without_repeating_characters

// Task: https://leetcode.com/problems/longest-substring-without-repeating-characters/

// Solution: https://leetcode.com/problems/longest-substring-without-repeating-characters/submissions/1041314028/

// tags:
// #Hash Table
// #String
// #Sliding Window

func lengthOfLongestSubstring(s string) int {
	used := make(map[byte]int)

	var longest, start int
	for i := 0; i < len(s); i++ {
		c := s[i]

		if j, ok := used[c]; ok && j >= start {
			start = used[c] + 1
		}

		longest = max(longest, i-start+1)

		used[c] = i
	}

	return longest
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
