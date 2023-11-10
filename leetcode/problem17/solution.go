package problem17

// Task: https://leetcode.com/problems/letter-combinations-of-a-phone-number/

// Solution: https://leetcode.com/problems/letter-combinations-of-a-phone-number/submissions/1096100151/

// tags:
// #String

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	prevCombinations := getAllCharsByDigit(digits[0])
	for i := 1; i < len(digits); i++ {
		newCombinations := make([]string, 0)
		for _, s := range prevCombinations {
			for _, c := range getAllCharsByDigit(digits[i]) {
				newCombinations = append(newCombinations, s+c)
			}
		}
		prevCombinations = newCombinations
	}

	return prevCombinations
}

func getAllCharsByDigit(b byte) []string {
	switch b {
	case '2':
		return []string{"a", "b", "c"}
	case '3':
		return []string{"d", "e", "f"}
	case '4':
		return []string{"g", "h", "i"}
	case '5':
		return []string{"j", "k", "l"}
	case '6':
		return []string{"m", "n", "o"}
	case '7':
		return []string{"p", "q", "r", "s"}
	case '8':
		return []string{"t", "u", "v"}
	case '9':
		return []string{"w", "x", "y", "z"}
	default:
		return nil
	}
}
