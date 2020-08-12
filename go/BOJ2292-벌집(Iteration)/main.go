package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
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

func NextInt(r *bufio.Scanner) int {
	r.Scan()
	ret, _ := strconv.Atoi(r.Text())

	return ret	
}

func NextStr(r *bufio.Scanner) string {
	r.Scan()
	return r.Text()
}

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)

	N := NextInt(r)
	var cnt int
	
	if N > 1 {
		for i := 1; i < N; i += cnt * 6 {
			cnt ++
		}
	}
	
	fmt.Fprintln(w, cnt + 1)
	
	w.Flush()
}