package problem67

// Task: https://leetcode.com/problems/add-binary/

// Solution: https://leetcode.com/problems/add-binary/submissions/1107055484/

// tags:
// #String
// #Math

// Time: O(max(N, M))
// Space: O(1)
func addBinary(a string, b string) string {
	result := make([]byte, 0)

	var hasReminder bool
	for iA, iB := len(a)-1, len(b)-1; iA >= 0 || iB >= 0; iA, iB = iA-1, iB-1 {
		var valA, valB, resVal byte = '0', '0', '0'
		if iA >= 0 {
			valA = a[iA]
		}
		if iB >= 0 {
			valB = b[iB]
		}

		if hasReminder {
			switch {
			case valA == '1' && valB == '1':
				resVal = '1'
				hasReminder = true
			case valA == '1' || valB == '1':
				resVal = '0'
				hasReminder = true
			default:
				resVal = '1'
				hasReminder = false
			}
		} else {
			switch {
			case valA == '1' && valB == '1':
				resVal = '0'
				hasReminder = true
			case valA == '1' || valB == '1':
				resVal = '1'
				hasReminder = false
			default:
				resVal = '0'
				hasReminder = false
			}
		}

		result = append(result, resVal)
	}

	if hasReminder {
		result = append(result, '1')
	}

	for l, r := 0, len(result)-1; l < r; l, r = l+1, r-1 {
		result[l], result[r] = result[r], result[l]
	}

	return string(result)
}
