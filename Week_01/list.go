package list

import "fmt"

//SliceList slice list
type SliceList struct {
	list []int
}

//Add element
func (sl *SliceList) Add(e int) {
	sl.list = append(sl.list, e)
}

//Insert element
func (sl *SliceList) Insert(index uint, e int) error {
	if int(index) > len(sl.list) {
		return fmt.Errorf("out of index")
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
