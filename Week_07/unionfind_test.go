package week7

import "testing"

func TestUnionFind(t *testing.T) {
	u := NewUF(10)
	u.union(1, 2)
	u.union(2, 3)
	u.union(3, 4)
	p1 := u.find(1)
	// p2 := u.find(2)
	// p3 := u.find(3)
	// p4 := u.find(4)
	t.Log(p1)
	t.Log(u)
	t.Log(u.count)

}
