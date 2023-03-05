package min_stack

type MinStack struct {
	stack         []int
	minVals       []int
	minValToCount map[int]int
}

func Constructor() MinStack {
	return MinStack{
		stack:         make([]int, 0),
		minVals:       make([]int, 0),
		minValToCount: make(map[int]int),
	}
}

func (m *MinStack) Push(val int) {
	if m == nil {
		return
	}

	stackWasEmpty := len(m.stack) == 0

	m.stack = append(m.stack, val)

	if stackWasEmpty {
		m.minVals = append(m.minVals, val)
		m.minValToCount[val]++
		return
	}

	currMinVal := m.minVals[len(m.minVals)-1]
	if hasNewMinVal := val < currMinVal; hasNewMinVal {
		m.minVals = append(m.minVals, val)
		m.minValToCount[val]++
		return
	}

	if thisOneOfExistsMinVals := m.minValToCount[val] > 0; thisOneOfExistsMinVals {
		m.minValToCount[val]++
		return
	}
}

func (m *MinStack) Pop() {
	if m == nil {
		return
	}

	if len(m.stack) == 0 {
		return
	}

	currLast, currMin := m.stack[len(m.stack)-1], m.minVals[len(m.minVals)-1]

	m.stack = m.stack[:len(m.stack)-1]

	if currLast == currMin {
		m.minValToCount[currMin]--
		if m.minValToCount[currMin] == 0 {
			m.minVals = m.minVals[:len(m.minVals)-1]
		}
		return
	}

	if currLastIsOneOfExistsMinVals := m.minValToCount[currLast] > 0; currLastIsOneOfExistsMinVals {
		m.minValToCount[currLast]--
		return
	}
}

func (m *MinStack) Top() int {
	if m == nil {
		return 0
	}

	if len(m.stack) == 0 {
		return 0
	}

	return m.stack[len(m.stack)-1]
}

func (m *MinStack) GetMin() int {
	if m == nil {
		return 0
	}

	if len(m.stack) == 0 {
		return 0
	}

	return m.minVals[len(m.minVals)-1]
}
