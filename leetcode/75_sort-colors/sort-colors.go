package sort_colors

// Task: https://leetcode.com/problems/sort-colors/

// Solution: https://leetcode.com/problems/sort-colors/submissions/1044687099/

// tags:
// #Array
// #Two Pointers
// #Sorting

func sortColors(nums []int) {
	var arrCount [3]int
	for i := 0; i < len(nums); i++ {
		arrCount[nums[i]]++
	}

	for i, k := 0, 0; i < len(arrCount); i++ {
		for j := 0; j < arrCount[i]; j, k = j+1, k+1 {
			nums[k] = i
		}
	}
}
