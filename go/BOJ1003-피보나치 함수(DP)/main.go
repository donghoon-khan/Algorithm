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
		n, _ := strconv.Atoi(r.Text())
		
		if n == 0 {
			fmt.Fprintf(w, "1 0\n")
		} else if n == 1 {
			fmt.Fprintf(w, "0 1\n")
		} else {
			zero := make([]int, n + 1)
			one := make([]int, n + 1)
					
			for j := 0; j <= n; j ++ {
				if j == 0 {
					zero[j] = 1
				} else if j == 1 {
					one[j] = 1
				} else {
					zero[j] = zero[j - 1] + zero[j - 2]
					one[j] = one[j - 1] + one[j - 2]
				}
			}
			fmt.Fprintf(w, "%d %d\n", zero[n], one[n])
			
		}
	}
	w.Flush()
}