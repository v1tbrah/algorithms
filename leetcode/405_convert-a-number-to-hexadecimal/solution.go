package problem405

import "bytes"

// Task: https://leetcode.com/problems/convert-a-number-to-hexadecimal/

// Solution: https://leetcode.com/problems/convert-a-number-to-hexadecimal/submissions/1094372737/

// tags:
// #Math
// #Bit Manipulation

// Time: O(Log(16)X)
// Space: O(Log(16)X)
func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	hexChars := map[int]byte{
		0:  '0',
		1:  '1',
		2:  '2',
		3:  '3',
		4:  '4',
		5:  '5',
		6:  '6',
		7:  '7',
		8:  '8',
		9:  '9',
		10: 'a',
		11: 'b',
		12: 'c',
		13: 'd',
		14: 'e',
		15: 'f',
	}

	if num < 0 {
		num += 1 << 32
	}

	var b bytes.Buffer
	for num > 0 {
		b.WriteByte(hexChars[num%16])
		num /= 16
	}

	result := b.Bytes()
	for l, r := 0, len(result)-1; l < r; l, r = l+1, r-1 {
		result[l], result[r] = result[r], result[l]
	}

	return string(result)
}
