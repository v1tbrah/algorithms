package string_compression

import (
	"strconv"
)

// Task: https://leetcode.com/problems/string-compression/

// Solution: https://leetcode.com/problems/string-compression/submissions/1066802011/

// tags:
// #String
// #Two Pointers

// Time: O(n)
// Space: O(1)
func compress(chars []byte) int {
	if len(chars) == 1 {
		return 1
	}

	startPrevCharIdx := 0
	prevChar := chars[0]
	lastUsedForSaveIdx := -1
	lastSavedChar := byte(126)
	for i := 1; i < len(chars); i++ {
		currChar := chars[i]
		if currChar != prevChar {
			charLen := i - startPrevCharIdx
			if charLen == 1 {
				chars[lastUsedForSaveIdx+1] = prevChar
				lastUsedForSaveIdx++
			} else {
				chars[lastUsedForSaveIdx+1] = prevChar
				lastUsedForSaveIdx++
				strCharLen := strconv.Itoa(charLen)
				for j := 0; j < len(strCharLen); j++ {
					chars[lastUsedForSaveIdx+1] = strCharLen[j]
					lastUsedForSaveIdx++
				}
			}

			lastSavedChar = prevChar

			prevChar = currChar

			startPrevCharIdx = i
		}
	}

	if lastSavedChar != chars[len(chars)-1] { // unprocessed last char
		charLen := len(chars) - startPrevCharIdx
		if charLen == 1 {
			chars[lastUsedForSaveIdx+1] = prevChar
			lastUsedForSaveIdx++
		} else {
			chars[lastUsedForSaveIdx+1] = prevChar
			lastUsedForSaveIdx++
			strCharLen := strconv.Itoa(charLen)
			for j := 0; j < len(strCharLen); j++ {
				chars[lastUsedForSaveIdx+1] = strCharLen[j]
				lastUsedForSaveIdx++
			}
		}
	}

	return lastUsedForSaveIdx + 1
}
