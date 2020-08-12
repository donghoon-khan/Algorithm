package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
	"math"
)

func min (a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	
	r.Scan()
	N, _ := strconv.Atoi(r.Text())
	dp := make([]int, N + 1)
	
	for i := 0; i <= N; i ++ {
		dp[i] = math.MaxInt32
	}
	
	dp[3] = 1
	if N >= 5 {
		dp[5] = 1
		for i := 6; i <= N; i ++ {
			if dp[i - 5] != math.MaxInt32 {
				dp[i] = dp[i - 5] + 1
			}
			if dp[i - 3] != math.MaxInt32 {
				dp[i] = min(dp[i], dp[i - 3] + 1)
			}
		}
	}
	
	if dp[N] == math.MaxInt32 {
		dp[N] = -1
	}	
	fmt.Fprintf(w, "%d\n", dp[N])
	w.Flush()
}