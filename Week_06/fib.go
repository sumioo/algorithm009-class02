package week6

import "fmt"

func fibDP(n uint) (uint, error) {
	dp := make([]uint, n)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1

	for i := uint(3); i < n; i++ {
		s := dp[i-1] + dp[i-2]
		if s > dp[i-1] && s > dp[i-2] {
			dp[i] = s
		} else {
			return 0, fmt.Errorf("overflow")
		}
	}
	return dp[n-1], nil
}

func fibDPFloat(n int) float64 {
	dp := make([]float64, n)
	dp[0] = 0.0
	dp[1] = 1.0

	for i := 3; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]

	}
	return dp[n-1]
}
