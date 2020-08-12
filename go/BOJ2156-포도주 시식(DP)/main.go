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
	arr := make([]int, N + 1)
	
	for i := 1; i <= N; i ++ {
		r.Scan()
		arr[i], _ = strconv.Atoi(r.Text())
	}
	
	dp := make([]int, N + 1)
	
	for i := 0; i <=N; i ++ {
		if i == 0 {
			dp[i] = 0
		} else if i == 1 {
			dp[i] = arr[i]
		} else if i == 2 {
			dp[i] = arr[i - 1] + arr[i]
		} else {
			dp[i] = max(max(dp[i - 2] + arr[i], dp[i - 3] + arr[i - 1] + arr[i]), dp[i - 1])
		}
	}
	
	

	fmt.Fprintf(w, "%d\n", dp[N])
	w.Flush()
}