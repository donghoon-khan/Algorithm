package main

import(
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
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)

	r.Scan()
	N, _ := strconv.Atoi(r.Text())
	step := make([]int, N + 1)
	
	for i := 1; i <= N; i ++ {
		r.Scan()
		step[i], _ = strconv.Atoi(r.Text())
	}
	
	dp := make([]int, N + 1)
	dp[0] = 0
	dp[1] = step[1]
	if (N >= 2) {
		dp[2] = step[1] + step[2]
	
		for i := 3; i <= N; i ++ {
			dp[i] = max(dp[i-2] + step[i], dp[i-3] + step[i - 1] + step[i])
		}
	}
	fmt.Fprintf(w, "%d\n", dp[N])
	w.Flush()
}