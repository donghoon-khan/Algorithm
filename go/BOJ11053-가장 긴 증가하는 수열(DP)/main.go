package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	
	w:= bufio.NewWriter(os.Stdout)
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)

	r.Scan()
	N, _ := strconv.Atoi(r.Text())
	
	arr := make([]int, N + 1)
	dp := make([]int, N + 1)
	
	result := 0
	for i := 1; i <= N; i ++ {
		r.Scan()
		arr[i], _ = strconv.Atoi(r.Text())
		dp[i] = 1
		for j := i - 1; j >= 0; j -- {
			if (arr[j] < arr[i]) {
				dp[i] = max (dp[i], dp[j] + 1)
			}
		}
		result = max(result, dp[i])
	}
	
	fmt.Fprintf(w, "%d\n", result)
	w.Flush()
}