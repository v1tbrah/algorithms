package problem11

// Task: https://leetcode.com/problems/container-with-most-water/

// Solution: https://leetcode.com/problems/container-with-most-water/submissions/1094595825/

// tags:
// #Array
// #Two Pointers

// Time: O(N)
// Space: O(1)
func maxArea(height []int) int {
	var maxVal int
	for l, r := 0, len(height)-1; l < r; {
		maxVal = max(maxVal, calcArea(height[l], height[r], r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return maxVal
}

func calcArea(h1, h2, w int) int {
	h := min(h1, h2)
	return h * w
}
