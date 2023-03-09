package daily_temperatures

func dailyTemperatures(temperatures []int) []int {

	res := make([]int, len(temperatures))

	tempInfo := make(map[int]int, 0) // idxInResult -> temp

	var prev int
	for i := 0; i < len(temperatures); i++ {

		tempInfo[i] = temperatures[i]

		if prev >= temperatures[i] {
			prev = temperatures[i]
			continue
		}

		for idxInResult, temp := range tempInfo {
			if idxInResult == i {
				continue
			}

			if temperatures[i] > temp {
				res[idxInResult] = i - idxInResult
				delete(tempInfo, idxInResult)
			}
		}

		prev = temperatures[i]
	}

	return res
}
