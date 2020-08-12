package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"sort"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	
	r.Scan()
	N, _ := strconv.Atoi(r.Text())
	
	num := make ([]int, N)
	for i := 0; i < N; i ++ {
		r.Scan()
		num[i], _ = strconv.Atoi(r.Text())
	}

	sort.Ints(num)

	r.Scan()
	M, _ := strconv.Atoi(r.Text())
	for i := 0; i < M; i ++ {
		r.Scan()
		in, _ := strconv.Atoi(r.Text())
		
		ret := sort.SearchInts(num, in)
		
		if ret == N {
			ret = 0
		} else {
			if num[ret] == in {
				ret = 1
			} else {
				ret = 0
			}
		}
		fmt.Fprintf(w, "%d\n", ret)
	}
	
	w.Flush()
	
}