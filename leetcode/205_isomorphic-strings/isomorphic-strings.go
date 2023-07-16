package isomorphic_strings

// Task: https://leetcode.com/problems/isomorphic-strings/

// Solution: https://leetcode.com/problems/isomorphic-strings/submissions/995736693/
// Thanks for this explanation: https://leetcode.com/problems/isomorphic-strings/solutions/1332570/my-solution-in-go-with-explanation/

// tags:
// #Hash Table
// #String

func isIsomorphic(s string, t string) bool {
	// s -> t
	byteMap := make(map[byte]byte)

	// t, who already used for someone s -> t
	alreadyMapped := make(map[byte]struct{})

	for i := 0; i < len(s); i++ {
		b, ok := byteMap[s[i]]
		if ok && b == t[i] { // if it's exists relation
			continue
		}

		if ok && b != t[i] { // if relation s -> t exists, but t[i] != t from this relation
			return false
		}

		if _, ok = alreadyMapped[t[i]]; ok { // if t[i] already mapped, but not for s[i]
			return false
		}

		byteMap[s[i]] = t[i]
		alreadyMapped[t[i]] = struct{}{}
	}

	return true
}
