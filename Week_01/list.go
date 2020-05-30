package list

import "fmt"
import "math"
import "sort"

//SliceList slice list
type SliceList struct {
	list []int
	sort.Sort
}

//Add element
func (sl *SliceList) Add(e int) {
	sl.list = append(sl.list, e)
}

//Insert element
func (sl *SliceList) Insert(index uint, e int) error {
	if int(index) > len(sl.list) {
		return fmt.Errorf("out of index")
		fmt.Fprint
	}

	sl.list = append(sl.list, 0)
	copy(sl.list[index+1:], sl.list[index:])
	sl.list[index] = e

	return nil
}

//Remove element
func (sl *SliceList) Remove(index uint) error {
	if int(index) >= len(sl.list) {
		return fmt.Errorf("out of index")
	}

	copy(sl.list[index:], sl.list[index+1:])
	sl.list = sl.list[:len(sl.list)-1]
	return nil
}

//IndexOf element
func (sl *SliceList) IndexOf(e int) int {
	for index, element := range sl.list {
		if element == e {
			return index
		}
	}
	return -1
}

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func (ll *LinkedList) Add(data int) *Node {
	node := &Node{data: data}
	if ll.tail == nil {
		ll.tail = node
		ll.head = node
	} else {
		ll.tail.next = node
	}
	ll.size++
	return node
}

func (ll *LinkedList) Remove(node *Node) {

}
