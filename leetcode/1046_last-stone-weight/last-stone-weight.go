package last_stone_weight

// Task: https://leetcode.com/problems/last-stone-weight/

// Solution: https://leetcode.com/problems/last-stone-weight/submissions/1054602166/

// tags:
// #Array
// #Heap (Priority Queue)

// Idea:
// 1. Make stones MaxHeap - O(n)
// 2. get 2 max elements 1 by 1 with heapify up-bottom on both steps - O(2*logN)
// 3. if 2 max is equal, continue - O(1)
// 4. if 2 max is not equal, get diff, push it to end, heapify bottom-up - O(logN)

// Time: O(n*logN)
// Space: O(logN) - logN - максимальный размер стека при рекурсии
func lastStoneWeight(stones []int) int {
	lastNonLeafIdx := len(stones)/2 - 1
	for i := lastNonLeafIdx; i >= 0; i-- {
		heapify(stones, i)
	}

	for len(stones) > 1 {
		last2 := [2]int{}
		for i := 0; i < 2; i++ {
			last2[i] = stones[0]
			stones[0] = stones[len(stones)-1]
			stones = stones[:len(stones)-1]
			heapify(stones, 0)
		}

		if last2[0] == last2[1] {
			continue
		}

		newEl := last2[0] - last2[1]
		if newEl < 0 {
			newEl = last2[1] - last2[0]
		}

		stones = append(stones, newEl)

		idx := len(stones) - 1
		parentIdx := (idx - 1) / 2

		for parentIdx >= 0 && stones[idx] > stones[parentIdx] {
			stones[idx], stones[parentIdx] = stones[parentIdx], stones[idx]
			idx = parentIdx
			parentIdx = (idx - 1) / 2
		}
	}

	if len(stones) == 0 {
		return 0
	}
	return stones[0]
}

func heapify(arr []int, idx int) {
	gratestValIdx := idx
	leftChildNodeIdx := idx*2 + 1
	rightChildNodeIdx := idx*2 + 2

	if leftChildNodeIdx < len(arr) && arr[gratestValIdx] < arr[leftChildNodeIdx] {
		gratestValIdx = leftChildNodeIdx
	}
	if rightChildNodeIdx < len(arr) && arr[gratestValIdx] < arr[rightChildNodeIdx] {
		gratestValIdx = rightChildNodeIdx
	}

	if idx != gratestValIdx {
		arr[idx], arr[gratestValIdx] = arr[gratestValIdx], arr[idx]
		heapify(arr, gratestValIdx)
	}
}
