package week7

type UF struct {
	count int
	id    []int
}

func NewUF(n int) *UF {
	x := make([]int, n)
	for i := range x {
		x[i] = i
	}
	return &UF{id: x, count: n}
}

func (u *UF) find(p int) int {
	root := p
	for root != u.id[root] {
		root = u.id[root]
	}

	for p != u.id[p] {
		p, u.id[p] = u.id[p], root
	}

	return root
}

func (u *UF) union(p int, q int) {
	r1 := u.find(p)
	r2 := u.find(q)
	if r1 == r2 {
		return
	}
	u.id[r1] = r2
	u.count--
}
