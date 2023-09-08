package sort

// Сортировка выбором
// Time: O(n^2)
// Space: O(1)
func selectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		minElIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minElIdx] {
				minElIdx = j
			}
		}
		arr[i], arr[minElIdx] = arr[minElIdx], arr[i]
	}
}
