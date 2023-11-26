package problem70

// Task: https://leetcode.com/problems/climbing-stairs/

// Solution: https://leetcode.com/problems/climbing-stairs/submissions/777244435/

// tags:
// #Math

// Time: O(N)
// Space: O(1)
func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	prevPrev, prev := 1, 2
	var curr int
	for i := 3; i <= n; i++ {
		curr = prevPrev + prev
		prevPrev, prev = prev, curr
	}
	return curr
}
