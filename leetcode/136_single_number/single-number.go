package single_number

// Task: https://leetcode.com/problems/single-number/

// Solution: https://leetcode.com/problems/single-number/submissions/993783517/

// tags:
// #Array
// #Bit Manipulation

func singleNumber(nums []int) int {
	numMap := make(map[int]int, len(nums))
	for _, num := range nums {
		if numMap[num] == 1 {
			delete(numMap, num)
		} else {
			numMap[num] = 1
		}
	}

	for num := range numMap { // numMap guarantee has only one neccessary element
		return num
	}

	return -1 // only for compile
}
