package sort

// Сортировка подсчетом
// Time: O(n+k), where k - is count of different elements
// Space: O(max)
func countingSort(nums []int) {
	if len(nums) == 0 {
		return
	}

	// find max element
	maxEl := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxEl {
			maxEl = nums[i]
		}
	}

	// make slice with length=maxEl+1
	countEls := make([]int, maxEl+1)

	// store the count of each element
	for i := 0; i < len(nums); i++ {
		countEls[nums[i]]++
	}

	// store the cumulative count of each
	for i := 1; i < len(countEls); i++ {
		countEls[i] += countEls[i-1]
	}
	// cumulative: [0,1,3,5,6,6,6,6,7]

	// find the index of each element of the original array in count array,
	// and place the elements in output array
	sortedNums := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		el := nums[i]

		idxInCountEls := countEls[el]
		sortedNums[idxInCountEls-1] = el

		// after used, decrease count
		countEls[el]--
	}

	// copy the sorted elements into original array
	for i := 0; i < len(nums); i++ {
		nums[i] = sortedNums[i]
	}
}
