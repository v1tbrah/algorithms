package k_closest_points_to_origin

// Task: https://leetcode.com/problems/k-closest-points-to-origin/

// Solution: https://leetcode.com/problems/k-closest-points-to-origin/submissions/1058928990/

// tags:
// #Array
// #Math
// #Heap (Priority Queue)
// #Geometry

type pointAndDistanceToStart struct {
	distance int
	point    []int
}

// Time: O(n+k*logn)
// Space: O(n)
func kClosest(points [][]int, k int) [][]int {
	// 1. make distances (square distances, cause we don't need to calculate sqrt) - TimeC O(n)
	// 2. make minHeap - TimeC O(n)
	// 3. get topKth min - TimeC O(klogn)

	// full TimeC - O(n+k*logn)

	// 1. make distances (square distances)
	pointsAndDistances := make([]pointAndDistanceToStart, len(points))
	for i, p := range points {
		pointsAndDistances[i] = calculateDistance(p)
	}

	// 2. make minHeap
	lastNonLeastIdx := (len(pointsAndDistances) - 1) / 2
	for i := lastNonLeastIdx; i >= 0; i-- {
		heapify(pointsAndDistances, i)
	}

	// 3. get topKth min
	result := make([][]int, 0, k)
	for i := 0; i < k; i++ {
		result = append(result, pointsAndDistances[0].point)

		pointsAndDistances[0] = pointsAndDistances[len(pointsAndDistances)-1]
		pointsAndDistances = pointsAndDistances[:len(pointsAndDistances)-1]
		heapify(pointsAndDistances, 0)
	}

	return result
}

func calculateDistance(point []int) pointAndDistanceToStart {
	// (x1 - x2)^2 + (y1 - y2)^2
	// 1) x1^2 - 2*x1*x2 + x2^2
	// 2) y1^2 - 2*y1*y2 + y2^2
	// 3) 1)+2)
	x1, x2 := point[0], 0
	y1, y2 := point[1], 0

	first := (x1 * x1) - 2*x1*x2 + (x2 * x2)
	second := (y1 * y1) - 2*y1*y2 + (y2 * y2)

	return pointAndDistanceToStart{
		distance: first + second,
		point:    point,
	}
}

func heapify(arr []pointAndDistanceToStart, idx int) {
	smallestValIdx := idx
	leftChildNodeIdx := idx*2 + 1
	rightChildNodeIdx := idx*2 + 2

	if leftChildNodeIdx < len(arr) && arr[smallestValIdx].distance > arr[leftChildNodeIdx].distance {
		smallestValIdx = leftChildNodeIdx
	}
	if rightChildNodeIdx < len(arr) && arr[smallestValIdx].distance > arr[rightChildNodeIdx].distance {
		smallestValIdx = rightChildNodeIdx
	}

	if idx != smallestValIdx {
		arr[idx], arr[smallestValIdx] = arr[smallestValIdx], arr[idx]
		heapify(arr, smallestValIdx)
	}
}
