package week2

import "testing"

func TestBinaryHeapInsert(t *testing.T) {
	heap := new(BinaryHeap)
	heap.init(5)
	heap.insert(5)
	heap.insert(3)
	heap.insert(30)
	heap.insert(4)
	heap.insert(12)
	t.Log(heap)
}

func TestBinaryHeapDeleteMax(t *testing.T) {
	heap := new(BinaryHeap)
	heap.init(5)
	heap.insert(5)
	heap.insert(3)
	heap.insert(30)
	heap.insert(4)
	heap.insert(12)
	m, _ := heap.deleteMax()
	t.Log(m, heap)
	m, _ = heap.deleteMax()
	t.Log(m, heap)
	m, _ = heap.deleteMax()
	t.Log(m, heap)
	m, _ = heap.deleteMax()
	t.Log(m, heap)
	m, _ = heap.deleteMax()
	t.Log(m, heap)
	heap.insert(120)
	t.Log(heap)
}
