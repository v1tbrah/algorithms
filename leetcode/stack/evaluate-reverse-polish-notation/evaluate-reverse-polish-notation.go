package evaluate_reverse_polish_notation

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for i := 0; i < len(tokens); i++ {
		operand, err := strconv.Atoi(tokens[i])
		if err == nil {
			stack = append(stack, operand)
			continue
		}

		if len(stack) < 2 {
			panic("unexpected input data")
		}

		operand1, operand2 := stack[len(stack)-2], stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		operationStr := tokens[i]
		operationResult := evaluateOperation(operationStr, operand1, operand2)
		stack = append(stack, operationResult)
	}

	if len(stack) != 1 {
		panic("unexpected input data")
	}
	return stack[0]
}

func evaluateOperation(operationStr string, operand1, operand2 int) int {
	switch operationStr {
	case "+":
		return operand1 + operand2
	case "-":
		return operand1 - operand2
	case "*":
		return operand1 * operand2
	case "/":
		return operand1 / operand2
	default:
		panic("unexpected input data")
	}
}
