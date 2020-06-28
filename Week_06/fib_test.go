package week6

import (
	"testing"
)

func TestFibDP(t *testing.T) {
	r, err := fibDP(94)
	t.Log("18446744073709551615")
	t.Log(r, err)

	rr := fibDPFloat(94)
	t.Log(rr)
}

// 1779979416004714189 90
// 18446744073709551615
// 12200160415121876738
