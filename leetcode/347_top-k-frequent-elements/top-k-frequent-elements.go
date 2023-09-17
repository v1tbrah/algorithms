package top_k_frequent_elements

// Task: https://leetcode.com/problems/top-k-frequent-elements/

// Solution: https://leetcode.com/problems/top-k-frequent-elements/submissions/1051845185/

// tags:
// #Array
// #Hash Table
// #Sorting
// #Heap (Priority Queue)

// Time: O(n+k*logN)
// n - for heapify input array; k*logN for get k elements from Heap and heapify after all get operations
// Space: O(n)
func topKFrequent(arr []int, k int) []int {
	// build map digit to frequency
	dfm := make(map[int]int)
	for _, d := range arr {
		dfm[d]++
	}

	// build map frequency to digits
	fdsm := make(map[int][]int)
	// build array frequencies
	fa := make([]int, 0)
	for d, f := range dfm {
		fa = append(fa, f)
		fdsm[f] = append(fdsm[f], d)
	}

	// build Max Heap frequencies
	lnoncei := len(fa)/2 - 1
	for i := lnoncei; i >= 0; i-- {
		heapify(fa, len(fa), i)
	}

	// move top k frequencies to end of Max Heap
	for i := 0; i < k; i++ {
		s := len(fa) - i
		fa[0], fa[s-1] = fa[s-1], fa[0]
		heapify(fa, s-1, 0)
	}

	// get top k frequencies
	topkf := fa[len(fa)-k:]

	// build result
	result := make([]int, 0, k)
	deduplicator := make(map[int]struct{})
	for _, f := range topkf {
		ds := fdsm[f]
		for _, d := range ds {
			if _, ok := deduplicator[d]; !ok {
				result = append(result, d)
				deduplicator[d] = struct{}{}
			}
		}
	}

	return result
}

func heapify(arr []int, s, i int) {
	lei := i
	lcni := i*2 + 1
	rcni := i*2 + 2

	if lcni < s && arr[lei] < arr[lcni] {
		lei = lcni
	}
	if rcni < s && arr[lei] < arr[rcni] {
		lei = rcni
	}

	if i != lei {
		arr[i], arr[lei] = arr[lei], arr[i]
		heapify(arr, s, lei)
	}
}
