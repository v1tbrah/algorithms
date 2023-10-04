package max_consecutive_ones_ii

// Task: https://leetcode.com/problems/max-consecutive-ones-ii/

// Solution: https://leetcode.com/problems/max-consecutive-ones-ii/submissions/1066898277/

// tags:
// #Array
// #Sliding Window

// Time: O(n)
// Space: O(1)
func findMaxConsecutiveOnes(nums []int) int {
	var wasFlip bool
	var maxCons int
	var withoutFlip int
	var withFlip int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			withoutFlip++
			withFlip++
			continue
		}

		wasFlip = true
		if withFlip > maxCons {
			maxCons = withFlip
		}
		withFlip = withoutFlip
		withoutFlip = 0
	}

	if withFlip > maxCons {
		maxCons = withFlip
	}

	if wasFlip {
		maxCons++
	}

	return maxCons
}
