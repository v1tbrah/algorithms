package happy_number

// Task: https://leetcode.com/problems/happy-number/

// Solution: https://leetcode.com/problems/happy-number/submissions/995142609/

// tags:
// #Hash Table
// #Math
// #Two Pointers

func isHappy(n int) bool {
	checked := make(map[int]struct{})
	for {
		var sum int
		for n > 0 {
			digit := n % 10
			sum += digit * digit
			n /= 10
		}

		if sum == 1 {
			return true
		}

		if _, ok := checked[sum]; ok {
			return false
		}

		checked[sum] = struct{}{}

		n = sum
	}
}
