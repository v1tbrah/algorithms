package design_circular_queue

type MyCircularQueue struct {
	queue   []int
	size    int
	busy    int
	headIdx int
	tailIdx int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		queue:   make([]int, k),
		size:    k,
		busy:    0,
		headIdx: -1,
		tailIdx: -1,
	}
}

func (q *MyCircularQueue) EnQueue(value int) bool {
	if q.IsFull() {
		return false
	}

	enqElIdx := q.tailIdx + 1
	if tailNowAtTheEnd := q.tailIdx == q.size-1; tailNowAtTheEnd {
		enqElIdx = 0
	}
	q.queue[enqElIdx] = value

	q.tailIdx = enqElIdx

	q.busy++

	if q.headIdx == -1 {
		q.headIdx = 0
	}

	return true
}

func (q *MyCircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}

	q.queue[q.headIdx] = 0

	newHeadIdx := q.headIdx + 1
	if headNowAtTheEnd := q.headIdx == q.size-1; headNowAtTheEnd {
		newHeadIdx = 0
	}
	q.headIdx = newHeadIdx

	q.busy--

	return true
}

func (q *MyCircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}

	return q.queue[q.headIdx]
}

func (q *MyCircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}

	return q.queue[q.tailIdx]
}

func (q *MyCircularQueue) IsEmpty() bool {
	return q.busy == 0
}

func (q *MyCircularQueue) IsFull() bool {
	return q.busy == q.size
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
