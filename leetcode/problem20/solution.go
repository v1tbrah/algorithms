package problem20

// Task: https://leetcode.com/problems/valid-parentheses/

// Solution: https://leetcode.com/problems/valid-parentheses/submissions/1067788912/

// tags:
// #String
// #Stack

// Time: O(n)
// Space: O(n)
func isValid(s string) bool {
	stackOpen := make([]rune, 0)
	for _, char := range s {
		if isClose(char) {
			if len(stackOpen) == 0 {
				// "}"
				return false // nothing to close
			}
			if !isCloseForOpen(char, stackOpen[len(stackOpen)-1]) {
				// {([})}
				return false // invalid close
			}
			stackOpen = stackOpen[:len(stackOpen)-1] // cut last
		} else {
			// open
			stackOpen = append(stackOpen, char)
		}
	}

	if len(stackOpen) > 0 {
		// if there is non closed
		return false
	}

	return true
}

func isClose(char rune) bool {
	return char == ')' ||
		char == ']' ||
		char == '}'

}

func isCloseForOpen(close, open rune) bool {
	return close == ')' && open == '(' ||
		close == ']' && open == '[' ||
		close == '}' && open == '{'
}
