package valid_parentheses

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if !isOpen(s[0]) {
		return false
	}

	openStack := []byte{s[0]}

	for i := 1; i < len(s); i++ {

		prev := s[i-1]
		curr := s[i]

		switch {

		case isOpen(prev) && isOpen(curr):
			openStack = append(openStack, curr)

		case isOpen(prev) && isClose(curr):
			if isValidOpenToClose(prev, curr) {
				openStack = openStack[:len(openStack)-1]
			} else {
				return false
			}

		case isClose(prev) && isClose(curr):
			if len(openStack) == 0 {
				return false
			}

			if isValidOpenToClose(openStack[len(openStack)-1], curr) {
				openStack = openStack[:len(openStack)-1]
			} else {
				return false
			}

		case isClose(prev) && isOpen(curr):

			openStack = append(openStack, curr)

		}

	}

	return len(openStack) == 0
}

func isOpen(b byte) bool {
	return b == '(' ||
		b == '[' ||
		b == '{'
}

func isClose(b byte) bool {
	return b == ')' ||
		b == ']' ||
		b == '}'
}

func isValidOpenToClose(o, c byte) bool {
	return o == '(' && c == ')' ||
		o == '[' && c == ']' ||
		o == '{' && c == '}'
}
