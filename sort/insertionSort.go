package sort

// Сортировка вставками.
// Применима на практике, когда размер массива небольшой, либо когда массив "почти отсортирован".
// Time: O(n^2)
// Space: O(1)
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
}
