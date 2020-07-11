package week8

import (
	"strconv"
	"strings"
)

type dLLNode struct {
	key   int
	value int
	prev  *dLLNode
	next  *dLLNode
}

type dLinkedList struct {
	head *dLLNode
	tail *dLLNode
}

func newDLinkedList() dLinkedList {
	dummyHead := &dLLNode{}
	dummyTail := &dLLNode{}
	dummyHead.next = dummyTail
	dummyTail.prev = dummyHead
	return dLinkedList{head: dummyHead, tail: dummyTail}
}

func (dl *dLinkedList) insertBefore(new *dLLNode, node *dLLNode) *dLLNode {
	new.prev = node.prev
	new.next = node
	node.prev.next = new
	node.prev = new
	return new
}

func (dl *dLinkedList) append(key int, value int) *dLLNode {
	new := &dLLNode{key: key, value: value}
	return dl.insertBefore(new, dl.tail)
}

func (dl *dLinkedList) remove(node *dLLNode) {
	if node == nil {
		return
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	node.next = nil
	node.prev = nil
}

func (dl *dLinkedList) removeTail() {
	dl.remove(dl.getTail())
}

func (dl *dLinkedList) moveToTail(n *dLLNode) {
	dl.remove(n)
	dl.insertBefore(n, dl.tail)
}

func (dl *dLinkedList) moveToHead(n *dLLNode) {
	dl.remove(n)
	after := dl.head.next
	dl.insertBefore(n, after)
}

func (dl *dLinkedList) insertFirst(key int, value int) *dLLNode {
	new := &dLLNode{key: key, value: value}
	after := dl.head.next
	dl.insertBefore(new, after)
	return new
}

func (dl *dLinkedList) getHead() *dLLNode {
	if dl.head.next == dl.tail {
		return nil
	}
	return dl.head.next
}

func (dl *dLinkedList) getTail() *dLLNode {
	if dl.tail.prev == dl.head {
		return nil
	}
	return dl.tail.prev
}

func (dl *dLinkedList) print() string {
	p := dl.getHead()
	keys := []string{}
	for p != dl.tail {
		keys = append(keys, strconv.Itoa(p.key))
		p = p.next
	}

	return strings.Join(keys, "->")
}
