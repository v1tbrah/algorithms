package maximum_population_year

func maximumPopulation(logs [][]int) int {
	yearToPopulation := make(map[int]int, 100)

	for y := 1950; y <= 2050; y++ {
		for _, log := range logs {
			// birth - log[0]
			// death - log[1]
			if y >= log[0] && y < log[1] {
				yearToPopulation[y]++
			}
		}
	}

	minY, maxP := 2051, -1 // minY gt max possible by constraints, maxP lt min possible
	for y, c := range yearToPopulation {
		if c > maxP || (c >= maxP && y < minY) {
			minY = y
			maxP = c
		}
	}

	return minY
}
