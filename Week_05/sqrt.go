package week5

import "fmt"

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func mySqrt(x int) float64 {
	xx := float64(x)
	left, right := 0.0, xx

	for i := 0; i < 100; i++ {
		mid := (left + right) / 2
		fmt.Println(i, mid)
		y := mid*mid - xx
		if abs(y) <= 0.1 {
			return mid
		} else if y < 0 {
			left = mid
		} else {
			right = mid
		}
	}
	return -1.0
}
