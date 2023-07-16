package contains_duplicate_ii

// Task: https://leetcode.com/problems/contains-duplicate-ii/

// Solution: https://leetcode.com/problems/contains-duplicate-ii/submissions/996122776/

// tags:
// #Array
// #Hash Table
// #Sliding Window

func containsNearbyDuplicate(nums []int, k int) bool {
	numsMap := make(map[int]int)
	for i, num := range nums {
		j, ok := numsMap[num]
		if !ok || abs(j-i) > k {
			numsMap[num] = i
		} else {
			return true
		}
	}
	return false
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
