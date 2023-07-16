package search_in_rotated_sorted_array

// Task: https://leetcode.com/problems/first-unique-character-in-a-string/

// Solution: https://leetcode.com/problems/first-unique-character-in-a-string/submissions/996104806/

// tags:
// #Hash Table
// #String
// #Queue
// #Counting

func firstUniqChar(s string) int {
	sMap := make(map[byte]bool) // map where dublicates has value 'true'
	for i := 0; i < len(s); i++ {
		if _, ok := sMap[s[i]]; ok {
			sMap[s[i]] = true
		} else {
			sMap[s[i]] = false
		}
	}

	for i := 0; i < len(s); i++ {
		if isDuplicate, ok := sMap[s[i]]; ok && !isDuplicate {
			return i
		}
	}

	return -1
}
