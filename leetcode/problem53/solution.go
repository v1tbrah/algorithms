package problem53

// Task: https://leetcode.com/problems/maximum-subarray/

// Solution: https://leetcode.com/problems/maximum-subarray/submissions/1106355356/

// tags:
// #Array

// Time: O(N)
// Space: O(1)
func maxSubArray(nums []int) int {
	var left int
	prevSum := nums[0]
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		for prevSum < 0 && left < i {
			prevSum -= nums[left]
			left++
		}

		prevSum += nums[i]
		maxSum = max(maxSum, prevSum)
	}

	return maxSum
}
