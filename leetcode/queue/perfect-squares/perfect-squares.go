package perfect_squares

func numSquares(n int) int {
	pfSquares := make([]int, 0)

	for i := 1; i*i <= n; i++ {
		pfSquares = append(pfSquares, i*i)
	}

	type scenario struct {
		n, steps int
	}

	queue := []scenario{{n, 0}}

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		for _, square := range pfSquares {
			newN := top.n - square
			switch {
			case newN == 0:
				return top.steps + 1
			case newN > 0:
				queue = append(queue, scenario{newN, top.steps + 1})
			}
		}
	}

	return -1
}
