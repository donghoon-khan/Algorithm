package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"sort"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	
	r.Scan()
	N, _ := strconv.Atoi(r.Text())
	
	arr := make([]int, N)
	for i := 0; i < N; i ++ {
		r.Scan()
		arr[i], _ = strconv.Atoi(r.Text())
	}
	sort.Ints(arr)
	for i := 0; i < N; i ++ {
		fmt.Fprintf(w, "%d\n", arr[i])
	}
	w.Flush()
}