package week8

import "testing"

func TestClimbeStairsR(t *testing.T) {
	r := climbStairsR(5)
	t.Log(r)
}

func TestClimbeStairsBU(t *testing.T) {
	r := climbStairsBU(5)
	t.Log(r)
}

func TestClimbeStairsBUoP(t *testing.T) {
	r := climbStairsBUoP(5)
	t.Log(r)
}

func TestClimbeStairsPaths(t *testing.T) {
	r := climbStairsPaths(6, []int{1, 2, 3})
	for i := range r {
		t.Log(r[i])
	}

}

func TestClimbeStairsPathsLimitChoice(t *testing.T) {
	r := climbStairsPathsLimitChoice(6, []int{1, 2, 3})
	for i := range r {
		t.Log(r[i])
	}
}

func TestClimbeStairsLimitChoice(t *testing.T) {
	r := climbStairsLimitChoice(5, []int{1, 2, 3})
	t.Log(r)
}

func TestClimbeStairsLimitChoiceBU(t *testing.T) {
	r := climbStairsLimitChoiceBU(6)
	t.Log(r)
}
