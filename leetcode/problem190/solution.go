package problem190

// Task: https://leetcode.com/problems/reverse-bits/

// Solution: https://leetcode.com/problems/reverse-bits/submissions/1095965495/

// tags:
// #Bit Manipulation

// Time: O(1)
// Space: O(1)
func reverseBits(num uint32) uint32 {
	lMask, rMask := uint32(1<<0), uint32(1<<31)
	for l, r := 0, 31; l < r; l, r = l+1, r-1 {
		lBitIsZero, rBitIsZero := num&lMask == 0, num&rMask == 0
		if lBitIsZero {
			num &= ^rMask
		} else {
			num |= rMask
		}
		if rBitIsZero {
			num &= ^lMask
		} else {
			num |= lMask
		}

		lMask, rMask = lMask<<1, rMask>>1
	}

	return num
}
