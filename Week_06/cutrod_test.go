package week6

import "testing"

func TestCutRod(t *testing.T) {
	palns := []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	cutPlans := cutRod(palns, 9)
	t.Log(cutPlans)
}

func TestCutRodMethod(t *testing.T) {
	r := cutMethod(4)
	t.Log(len(r))
	t.Log(r)

}

func TestCutRodMethodB(t *testing.T) {
	r := cutMethodB(10)
	t.Log(len(r))
	// t.Log(r)

}

func TestCutRodMethodC(t *testing.T) {
	r := cutMethodC(4)
	t.Log(r)
	// t.Log(r)

}
