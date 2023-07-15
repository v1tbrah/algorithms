package two_sum

// Task: https://leetcode.com/problems/two-sum/

// Solution: https://leetcode.com/problems/two-sum/submissions/995378408/

// tags:
// #Array
// #Hash Table

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i, num := range nums {
		if j, ok := numsMap[target-num]; ok && j != i {
			return []int{i, j}
		}
		numsMap[num] = i
	}

	return nil
}
