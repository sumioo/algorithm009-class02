package week6

import "testing"

var matrix [][]int = [][]int{
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 1, 0, 0, 0, 1, 0},
	{0, 0, 0, 0, 1, 0, 0, 0},
	{1, 0, 1, 0, 0, 1, 0, 0},
	{0, 0, 1, 0, 0, 0, 0, 0},
	{0, 0, 0, 1, 1, 0, 1, 0},
	{0, 1, 0, 0, 0, 1, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
}

func TestCountPathsBacktrace(t *testing.T) {
	count := countPathsBacktrace(matrix)
	t.Log(count)
}

func TestCountPathsTopDown(t *testing.T) {
	count := countPathsDPTopDown(matrix)
	t.Log(count)
}

func TestCountPathsBottomUp(t *testing.T) {
	count := countPathsDPBottomUp(matrix)
	t.Log(count)
}
