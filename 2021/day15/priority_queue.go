package day15

// from https://gist.github.com/vderyagin/4573099

import (
	"errors"
)

type Elem struct {
	Score int
	Data  Point
}

type priorityQueue struct {
	heap []*Elem
}

func New(elems ...*Elem) *priorityQueue {
	pq := new(priorityQueue)

	pq.heap = elems

	for i := pq.lastNonLeafIdx(); i >= 0; i-- {
		pq.bubbleDown(i)
	}

	return pq
}

func (pq *priorityQueue) IsEmpty() bool {
	if len(pq.heap) == 0 {
		return true
	}

	return false
}

func (pq *priorityQueue) Peek() (*Elem, error) {
	if pq.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	return pq.heap[0], nil
}

func (pq *priorityQueue) Enqueue(e *Elem) {
	pq.heap = append(pq.heap, e)
	pq.bubbleUp(len(pq.heap) - 1)
}

func (pq *priorityQueue) Dequeue() (*Elem, error) {
	elem, err := pq.Peek()

	if err != nil {
		return nil, err
	}

	pq.replaceRootWithLastChild()
	pq.bubbleDown(0)

	pq.downsizeHeapIfTooBig()

	return elem, nil
}

func (pq *priorityQueue) parentIdx(idx int) int {
	return (idx - 1) / 2
}

func (pq *priorityQueue) childrenIdxs(idx int) []int {
	idxs := make([]int, 0, 2)

	xIdx := idx*2 + 1
	yIdx := idx*2 + 2

	if xIdx < len(pq.heap) {
		idxs = append(idxs, xIdx)
	}

	if yIdx < len(pq.heap) {
		idxs = append(idxs, yIdx)
	}

	return idxs
}

func (pq *priorityQueue) hasSmallerChild(idx int) bool {
	if idx >= len(pq.heap) {
		return false
	}

	idxs := pq.childrenIdxs(idx)
	score := pq.heap[idx].Score

	for _, i := range idxs {
		childScore := pq.heap[i].Score
		if score > childScore {
			return true
		}
	}

	return false
}

func (pq *priorityQueue) hasBiggerParent(idx int) bool {
	score := pq.heap[idx].Score
	parentScore := pq.heap[pq.parentIdx(idx)].Score

	return parentScore > score
}

func (pq *priorityQueue) bubbleUp(idx int) {
	if pq.hasBiggerParent(idx) {
		pq.bubbleUp(pq.switchWithParent(idx))
	} else {
		return
	}
}

func (pq *priorityQueue) bubbleDown(idx int) {
	if pq.hasSmallerChild(idx) {
		pq.bubbleDown(pq.switchWithSmallestChild(idx))
	} else {
		return
	}
}

func (pq *priorityQueue) switchWithParent(idx int) (newIdx int) {
	parentIdx := pq.parentIdx(idx)

	h := pq.heap
	h[parentIdx], h[idx] = h[idx], h[parentIdx]

	return parentIdx
}

func (pq *priorityQueue) switchWithSmallestChild(idx int) (newIdx int) {
	h := pq.heap
	newIdx = idx

	for _, i := range pq.childrenIdxs(idx) {
		if h[i].Score < h[newIdx].Score {
			newIdx = i
		}
	}

	h[idx], h[newIdx] = h[newIdx], h[idx]

	return newIdx
}

func (pq *priorityQueue) replaceRootWithLastChild() {
	lastIdx := len(pq.heap) - 1
	pq.heap[0] = pq.heap[lastIdx]
	pq.heap[lastIdx] = nil
	pq.heap = pq.heap[:lastIdx]
}

func (pq *priorityQueue) lastNonLeafIdx() int {
	lastIdx := len(pq.heap) - 1

	return pq.parentIdx(lastIdx)
}

func (pq *priorityQueue) downsizeHeapIfTooBig() {
	l := len(pq.heap)
	c := cap(pq.heap)

	if c > 100 && c > l*4 {
		newHeap := make([]*Elem, l, l*2)

		for idx, elem := range pq.heap {
			newHeap[idx] = elem
		}

		pq.heap = newHeap
	}
}
