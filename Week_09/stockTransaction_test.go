package week9

import "testing"

func TestStockTransaction(t *testing.T) {
	prices := []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}
	m := maxProfit122(prices)
	t.Log(m)
}
