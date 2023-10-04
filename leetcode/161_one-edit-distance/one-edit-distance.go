package one_edit_distance

import (
	"bytes"
)

// Task: https://leetcode.com/problems/one-edit-distance/

// Solution: https://leetcode.com/problems/one-edit-distance/submissions/1066733933/

// tags:
// #String
// #Two Pointers

// Time: O(n)
// Space: O(1)
func isOneEditDistance(s string, t string) bool {
	small, big := s, t
	if len(s) > len(t) {
		small, big = t, s
	}

	if len(big)-len(small) > 1 {
		return false
	}

	for i, j := 0, 0; i < len(small) && j < len(big); i, j = i+1, j+1 {
		if small[i] != big[j] {
			if len(small) != len(big) {
				return bytes.Equal([]byte(small[i:]), []byte(big[i+1:])) // DELETE scenario
			} else {
				return bytes.Equal([]byte(small[i+1:]), []byte(big[i+1:])) // REPLACE scenario
			}
		}
	}

	return len(big)-len(small) == 1 // INSERT scenario
}
