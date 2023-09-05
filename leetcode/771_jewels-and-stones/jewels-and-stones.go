package jewels_and_stones

// Task: https://leetcode.com/problems/jewels-and-stones/

// Solution: https://leetcode.com/problems/jewels-and-stones/submissions/1041282989/

// tags:
// #Hash Table
// #String

func numJewelsInStones(jewels string, stones string) int {
	jewelsMap := make(map[byte]struct{}, len(jewels))
	for i := 0; i < len(jewels); i++ {
		jewelsMap[jewels[i]] = struct{}{}
	}

	var sum int
	for i := 0; i < len(stones); i++ {
		if _, ok := jewelsMap[stones[i]]; ok {
			sum++
		}
	}

	return sum
}
