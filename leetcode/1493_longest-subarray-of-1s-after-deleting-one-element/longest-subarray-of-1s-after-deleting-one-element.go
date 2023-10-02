package longest_subarray_of_1s_after_deleting_one_element

// Task: https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/

// Solution: https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/submissions/1065192210/

// tags:
// #Array
// #Sliding Window

// Time: O(n)
// Space: O(1)
func longestSubarray(nums []int) int {
	var max, withOutDel, withDel int
	var isDeletedAtLestOnce bool
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			isDeletedAtLestOnce = true
			if withDel > max {
				max = withDel
			}
			withDel = withOutDel
			withOutDel = 0
			continue
		}

		withDel++
		withOutDel++
	}

	if !isDeletedAtLestOnce {
		return len(nums) - 1
	}

	if withDel > max {
		max = withDel
	}

	return max
}
