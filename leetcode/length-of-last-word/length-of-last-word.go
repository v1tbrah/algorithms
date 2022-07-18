package length_of_last_word

import (
	"strings"
)

// Implementation with package strings
func LengthOfLastWordV1(s string) int {
	s = strings.TrimRight(s, " ")
	sep := strings.Split(s, " ")
	return len(sep[len(sep)-1])
}

// Implementation without package strings
func LengthOfLastWordV2(s string) int {
	result := 0
	startLastWord := 0
	prevIsSpace := false
	for i, v := range s {
		isSpace := v == ' '
		if isSpace {
			prevIsSpace = true
		} else {
			if prevIsSpace {
				startLastWord = i
				prevIsSpace = false
			}
			result = i - startLastWord + 1
		}
	}
	return result
}
