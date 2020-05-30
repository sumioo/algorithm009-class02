package week2

import "fmt"

//BinaryHeap heap data struct
type BinaryHeap struct {
	heap []int
	size int
	max  int
}

func (h *BinaryHeap) init(max int) {
	h.heap = make([]int, max+1)
	h.max = max
}

func (h *BinaryHeap) String() string {
	return fmt.Sprintf("{heap:%v size:%d max:%d", h.heap[1:h.size+1], h.size, h.max)
}

func (h *BinaryHeap) deleteMax() (int, error) {
	if h.size == 0 {
		return -1, fmt.Errorf("empty heap")
	}
	max := h.heap[1]
	h.heap[1] = h.heap[h.size]
	h.size--
	h.sink(1)
	return max, nil
}

func (h *BinaryHeap) insert(value int) error {
	if h.size == h.max {
		return fmt.Errorf("full heap")
	}
	h.heap[h.size+1] = value
	h.size++
	h.swim(h.size)
	return nil
}

func (h *BinaryHeap) less(i int, j int) bool {
	if h.heap[i] < h.heap[j] {
		return true
	}
	return false
}

func (h *BinaryHeap) swap(i int, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

func (h *BinaryHeap) swim(k int) {
	for k > 1 && h.less(k/2, k) {
		h.swap(k/2, k)
		k = k / 2
	}
}

func (h *BinaryHeap) sink(k int) {
	for 2*k <= h.size {
		i := 2 * k
		if i < h.size && h.less(i, i+1) {
			i++
		}
		if !h.less(k, i) {
			break
		}
		h.swap(k, i)
		k = i
	}
}
