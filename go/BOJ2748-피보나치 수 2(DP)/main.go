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
	
	fibonacci := make([]int, N + 1)
	fibonacci[0] = 0
	fibonacci[1] = 1
	
	for i := 2; i <= N; i ++ {
		fibonacci[i] = fibonacci[i - 2] + fibonacci[i - 1]
	}
	
	fmt.Fprintf(w, "%d\n", fibonacci[N])
	w.Flush()
}