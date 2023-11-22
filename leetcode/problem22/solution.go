package problem22

// Task: https://leetcode.com/problems/generate-parentheses/

// Solution: https://leetcode.com/problems/generate-parentheses/submissions/1104447034/

// tags:
// #String

func generateParenthesis(n int) []string {
	type tmp struct {
		data  []byte
		open  int
		close int
	}

	currLvl := []tmp{{data: []byte{'('}, open: 1}}

	result := make([]string, 0)

	for len(currLvl) > 0 {
		newLvl := make([]tmp, 0)
		for _, info := range currLvl {
			if len(info.data) == n*2 {
				result = append(result, string(info.data))
				continue
			}

			if info.open < n {
				newInfo := tmp{
					open:  info.open + 1,
					close: info.close,
					data:  make([]byte, len(info.data)+1),
				}
				copy(newInfo.data, info.data)
				newInfo.data[len(newInfo.data)-1] = '('
				newLvl = append(newLvl, newInfo)
			}
			if info.open > info.close {
				newInfo := tmp{
					open:  info.open,
					close: info.close + 1,
					data:  make([]byte, len(info.data)+1),
				}
				copy(newInfo.data, info.data)
				newInfo.data[len(newInfo.data)-1] = ')'
				newLvl = append(newLvl, newInfo)
			}
		}

		currLvl = newLvl
	}

	return result
}
