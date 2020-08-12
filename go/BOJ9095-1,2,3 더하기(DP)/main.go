package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	
	r.Scan()
	N, _ := strconv.Atoi(r.Text())
	dp := make([]int, 11)
	
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	dp[3] = 4
	
	for i := 4; i < 11; i ++ {
		dp[i] = dp[i - 1] + dp[i - 2] + dp[i - 3]
	}
	
	for i := 0; i < N; i ++ {
		r.Scan()
		num, _ := strconv.Atoi(r.Text())
		fmt.Fprintf(w, "%d\n", dp[num])
	}
			
	w.Flush()
}