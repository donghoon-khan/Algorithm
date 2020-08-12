package main

import (
  "fmt" 
)

func min (a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

	var X int
	fmt.Scanln(&X)
	
	dp := make([]int, X + 1)
	
	dp[0] = 0
	dp[1] = 0
	
	for i := 2; i <= X; i ++ {
		dp[i] = dp[i - 1] + 1
		if i % 2 == 0 {
			dp[i] = min(dp[i], dp[i / 2] + 1)
		}
		if i % 3 == 0 {
			dp[i] = min(dp[i], dp[i / 3] + 1)
		}
	}
	
	fmt.Print(dp[X])
}