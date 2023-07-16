package group_anagrams

import "sort"

// Task: https://leetcode.com/problems/group-anagrams/

// Solution: https://leetcode.com/problems/group-anagrams/submissions/996133762/

// tags:
// #Array
// #Hash Table
// #String
// #Sorting

func groupAnagrams(strs []string) [][]string {
	mapAnagrams := make(map[string][]string)
	for _, str := range strs {
		srcBytes := []byte(str)
		sort.Slice(srcBytes, func(i, j int) bool {
			return srcBytes[i] <= srcBytes[j]
		})
		sortedStr := string(srcBytes)
		mapAnagrams[sortedStr] = append(mapAnagrams[sortedStr], str)
	}

	result := make([][]string, 0, len(mapAnagrams))
	for _, anagrams := range mapAnagrams {
		result = append(result, anagrams)
	}
	return result
}
