package week8

import "testing"

func TestTotalQueues(t *testing.T) {
	total := totalNQueens(4)
	t.Log(total)
}
