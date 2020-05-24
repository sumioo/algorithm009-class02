package list

import "testing"

func TestSliceListAdd(t *testing.T) {
	sl := SliceList{}
	sl.Add(0)
	sl.Add(1)
	sl.Add(2)
	t.Log(sl)
}

func TestSliceListInsert(t *testing.T) {
	sl := SliceList{list: []int{0, 1, 2, 3, 4, 5}}
	sl.Insert(0, 100)
	t.Log(sl)
	sl.Insert(2, 100)
	t.Log(sl)
	sl.Insert(3, 103)
	t.Log(sl)

	err := sl.Insert(20, 100)
	t.Log(err)
}

func TestSliceListRemove(t *testing.T) {
	sl := SliceList{list: []int{0, 1, 2, 3, 4, 5}}
	sl.Remove(0)
	t.Log(sl)
	sl.Remove(0)
	t.Log(sl)
	sl.Remove(2)
	t.Log(sl)
	sl.Remove(0)
	t.Log(sl)
	sl.Remove(0)
	t.Log(sl)
	sl.Remove(0)
	t.Log(sl)
	err := sl.Remove(0)
	t.Log(err)
}

func TestSliceListIndexOf(t *testing.T) {
	sl := SliceList{list: []int{0, 1, 2, 3, 4, 5}}
	index := sl.IndexOf(1)
	t.Log(index)
}
