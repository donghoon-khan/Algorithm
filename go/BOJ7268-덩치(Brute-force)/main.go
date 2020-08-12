package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
)

var (
	r *bufio.Scanner
	w *bufio.Writer
)

func Min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Init() {
	
	w = bufio.NewWriter(os.Stdout)
	r = bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords) 
}

func NextInt() int {
	r.Scan()
	ret, _ := strconv.Atoi(r.Text())

	return ret
}

func NextStr() string {
	r.Scan()
	return r.Text()
}

func main() {
	
	Init()
	
	N := NextInt()
	x := make([]int, N)
	y := make([]int, N)
	
	for i := 0; i < N; i ++ {
		x[i] = NextInt()
		y[i] = NextInt()
	}
	
	for i := 0; i < N; i ++ {
		cnt := 0
		for j := 0; j < N; j ++ {
			if i == j {
				continue
			} else if x[i] < x[j] && y[i] < y[j] {
				cnt ++
			}
		}
		fmt.Fprintf(w, "%d ", cnt + 1)
	}
	
	w.Flush()
}