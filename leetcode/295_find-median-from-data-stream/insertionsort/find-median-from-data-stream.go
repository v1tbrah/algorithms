package insertionsort

// Task: https://leetcode.com/problems/find-median-from-data-stream/

// Solution: // Task: https://leetcode.com/problems/find-median-from-data-stream/submissions/1063183203/

// tags:
// #Sorting
// #Data Stream
// #Design

type MedianFinder struct {
	arr []int
}

func Constructor() MedianFinder {
	return MedianFinder{
		arr: make([]int, 0),
	}
}

// Time: O(n)
// Space: O(1)
func (this *MedianFinder) AddNum(num int) {
	this.arr = append(this.arr, num)
	if len(this.arr) == 1 {
		return
	}

	// all times, when we go there, it sorted, exclude last element.
	// 	after push last element, we can sort it by 'insertion sort' by O(n) time
	for j := len(this.arr) - 1; j > 0; j-- {
		if this.arr[j-1] > this.arr[j] {
			this.arr[j-1], this.arr[j] = this.arr[j], this.arr[j-1]
		} else {
			break
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.arr) == 0 {
		return 0
	}

	if len(this.arr)%2 == 0 {
		return float64(this.arr[len(this.arr)/2-1]+this.arr[len(this.arr)/2]) / 2
	}

	return float64(this.arr[len(this.arr)/2])
}
