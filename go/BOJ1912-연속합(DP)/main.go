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
	
	dp := make([]int, N + 1)
	dp[0] = -1001
	result := -1001
	for i := 1; i <= N; i ++ {
		r.Scan()
		num, _ := strconv.Atoi(r.Text())
		
		dp[i] = max(dp[i-1] + num, num)
		result = max(result, dp[i])
		
	}
	
	fmt.Fprintf(w, "%d\n", result)
	w.Flush()
}