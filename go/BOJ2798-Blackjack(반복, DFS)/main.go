package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

var N, M int
var card [101]int
var result int

func findResult(depth int, idx int, sum int) {
	
	if depth == 3 {
		if M - sum >= 0 && result < sum {
			result = sum
		}
		return
	}
	
	for i := idx; i < N; i ++ {
		findResult(depth + 1, i + 1, sum + card[i])
	}
}

func main() {
	
	w:= bufio.NewWriter(os.Stdout)
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	
	r.Scan()
	N, _ = strconv.Atoi(r.Text())
	
	r.Scan()
	M, _ = strconv.Atoi(r.Text())
	
	for i := 0; i < N; i ++ {
		r.Scan()
		card[i], _ = strconv.Atoi(r.Text())
	}
	
	findResult(0, 0, 0)
	
	fmt.Println(result)
	
	w.Flush()
}