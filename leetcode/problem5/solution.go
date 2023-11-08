package problem5

// Task: https://leetcode.com/problems/longest-palindromic-substring/

// Solution: https://leetcode.com/problems/longest-palindromic-substring/submissions/1091582695/

// tags:
// #String

// Time: O(N^2)
// Space: O(1)
func longestPalindrome(s string) string {
	leftRes, rightRes := -1, -1
	for i := 0; i < len(s); i++ {
		if l, r := expandFromCenters(s, i, i); r-l >= rightRes-leftRes {
			leftRes = l
			rightRes = r
		}
		if l, r := expandFromCenters(s, i, i+1); r-l >= rightRes-leftRes {
			leftRes = l
			rightRes = r
		}
	}

	if leftRes == -1 && rightRes == -1 {
		return ""
	}

	return s[leftRes : rightRes+1]
}

func expandFromCenters(s string, l, r int) (int, int) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}

	return l + 1, r - 1
}
