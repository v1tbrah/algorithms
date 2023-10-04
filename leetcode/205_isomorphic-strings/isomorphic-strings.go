package isomorphic_strings

// Task: https://leetcode.com/problems/isomorphic-strings/

// Solution: https://leetcode.com/problems/isomorphic-strings/submissions/1067169271/

// tags:
// #Hash Table
// #String

// Time: O(n)
// Space: O(n)
func isIsomorphic(s string, t string) bool {
	mapped := make(map[byte]byte)
	usedB := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		a := s[i]
		b := t[i]

		mappedA, alreadyMapped := mapped[a]
		if !alreadyMapped {
			mapped[a] = b
			if usedB[b] {
				// we need map, but it already used in 't'
				// example: [e: a, a: d], 'a' already used
				return false
			}
			usedB[b] = true
		} else if mappedA != b {
			// we already mapped, but not for this char
			// example: [e: a, e: d], 'e' already mapped to 'a', we can't swap it to 'd'
			return false
		}
	}

	return true
}
