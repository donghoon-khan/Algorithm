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
	
	r.Scan()
	K, _ := strconv.Atoi(r.Text())
	
	coin := make([]int, N)
	for i := 0; i < N; i ++ {
		r.Scan()
		coin[i], _ = strconv.Atoi(r.Text())
	}
	
	var result int
	idx := N - 1
	for K > 0 {
		if coin[idx] <= K {
			K -= coin[idx]
			result++
		} else {
			idx --
		}
	}
	
	fmt.Fprintf(w, "%d\n", result)
	w.Flush()
}