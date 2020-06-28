package week6

import (
	"testing"
)

func printLCS(dp [][]int, i int, j int, s1 string, s2 string) string {
	if i == 0 || j == 0 {
		return ""
	}
	var s string
	if s1[i-1] == s2[j-1] {
		s = printLCS(dp, i-1, j-1, s1, s2)
		s = s + string(s1[i-1])
	} else if dp[i-1][j] >= dp[i][j-1] {
		s = printLCS(dp, i-1, j, s1, s2)
	} else {
		s = printLCS(dp, i, j-1, s1, s2)
	}
	return s
}
func TestLCS(t *testing.T) {
	s1 := "ABCBDAB"
	s2 := "BDCABAB"
	dp := lcs(s1, s2)
	t.Log(dp)
	s := printLCS(dp, len(dp)-1, len(dp[0])-1, s1, s2)
	t.Log(s)
}
