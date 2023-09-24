package intuition

import "sort"

// Task: https://leetcode.com/problems/meeting-rooms-ii/

// Solution: https://leetcode.com/problems/meeting-rooms-ii/submissions/1058106531/

// tags:
// #Array
// #Sorting
// #Heap (Priority Queue)

// Time: O(n*logN)
// Space: O(n)
func minMeetingRooms(intervals [][]int) int {
	// sort by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	meetingRoomsEnds := make([]int, 0)
	var maxRooms int
	for i := 0; i < len(intervals); i++ {
		currStart := intervals[i][0]

		// remove ends, that less than curr start
		// minEnd - meetingRoomsEnds[0]
		for len(meetingRoomsEnds) > 0 && meetingRoomsEnds[0] < currStart {
			meetingRoomsEnds[0] = meetingRoomsEnds[len(meetingRoomsEnds)-1]
			meetingRoomsEnds = meetingRoomsEnds[:len(meetingRoomsEnds)-1]
			heapifyUpToBottom(meetingRoomsEnds, 0)
		}

		currEnd := intervals[i][1]

		if len(meetingRoomsEnds) == 0 || currStart < meetingRoomsEnds[0] { // minEnd - meetingRoomsEnds[0]
			meetingRoomsEnds = append(meetingRoomsEnds, currEnd)
		} else if currStart == meetingRoomsEnds[0] { // minEnd - meetingRoomsEnds[0]
			meetingRoomsEnds[0] = currEnd
		}

		heapifyBottomToUp(meetingRoomsEnds)
		if maxRooms < len(meetingRoomsEnds) {
			maxRooms = len(meetingRoomsEnds)
		}
	}

	return maxRooms
}

func heapifyBottomToUp(arr []int) {
	idx := len(arr) - 1
	parentIdx := (idx - 1) / 2

	for parentIdx >= 0 && arr[idx] < arr[parentIdx] {
		arr[idx], arr[parentIdx] = arr[parentIdx], arr[idx]
		idx = parentIdx
		parentIdx = (idx - 1) / 2
	}
}

func heapifyUpToBottom(arr []int, idx int) {
	smallestValIdx := idx
	leftChildNodeIdx := idx*2 + 1
	rightChildNodeIdx := idx*2 + 2

	if leftChildNodeIdx < len(arr) && arr[smallestValIdx] > arr[leftChildNodeIdx] {
		smallestValIdx = leftChildNodeIdx
	}
	if rightChildNodeIdx < len(arr) && arr[smallestValIdx] > arr[rightChildNodeIdx] {
		smallestValIdx = rightChildNodeIdx
	}

	if idx != smallestValIdx {
		arr[idx], arr[smallestValIdx] = arr[smallestValIdx], arr[idx]
		heapifyUpToBottom(arr, smallestValIdx)
	}
}
