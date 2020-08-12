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
	tc, _ := strconv.Atoi(r.Text())
	
	for i := 0; i < tc; i ++ {
		r.Scan()
		str := r.Text()
		dp := make([]int, len(str) + 1)
		result := 0 
		for i := 1; i <= len(str); i ++ {
			if str[i - 1] == 'X' {
				dp[i] = 0
			} else {
				dp[i] = dp[i - 1] + 1
				result += dp[i]
			}
		}
		
		fmt.Fprintf(w, "%d\n", result)
	}
	w.Flush()
}