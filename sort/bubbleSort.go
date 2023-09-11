package sort

// Сортировка пузырьком
// Time: O(n^2)
// Space: O(1)
func bubbleSort(input []int) {
	for isSwapped := true; isSwapped; {
		isSwapped = false
		for i := 0; i < len(input)-1; i++ {
			if input[i+1] < input[i] {
				input[i], input[i+1] = input[i+1], input[i]
				isSwapped = true
			}
		}
	}
}
