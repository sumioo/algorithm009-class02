package week8

import "testing"

func TestDoubleLinkList(t *testing.T) {
	dl := newDLinkedList()
	dl.append(1, 1)
	dl.append(2, 1)
	dl.append(3, 1)
	dl.append(4, 1)
	t.Log(dl.print())
}

func TestDoubleLinkListMoveToTail(t *testing.T) {
	dl := newDLinkedList()
	dl.append(1, 1)
	dl.append(2, 1)
	dl.append(3, 1)
	dl.append(4, 1)
	t.Log(dl.print())
	n5 := dl.append(5, 1)
	dl.append(6, 1)
	t.Log(dl.print())
	dl.moveToTail(n5)
	t.Log(dl.print())
	dl.moveToTail(dl.getTail())
	t.Log(dl.print())
}

func TestDoubleLinkListRemoveTail(t *testing.T) {
	dl := newDLinkedList()
	dl.append(1, 1)
	dl.append(2, 1)
	dl.append(3, 1)
	dl.append(4, 1)
	dl.removeTail()
	t.Log(dl.print())
}

func TestDoubleLinkListMoveToHead(t *testing.T) {
	dl := newDLinkedList()
	dl.append(1, 1)
	dl.append(2, 1)
	dl.append(3, 1)
	n4 := dl.append(4, 1)
	t.Log(dl.print())
	dl.moveToHead(n4)
	t.Log(dl.print())
	dl.moveToHead(n4)
	t.Log(dl.print())
}

func TestDoubleLinkListInsertFirst(t *testing.T) {
	dl := newDLinkedList()
	dl.append(1, 1)
	dl.append(2, 1)
	dl.append(3, 1)
	dl.append(4, 1)
	t.Log(dl.print())
	dl.insertFirst(7, 1)
	t.Log(dl.print())
	dl.insertFirst(8, 1)
	t.Log(dl.print())
}
