package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	defer writer.Flush()

	var n int
	scanf("%d\n", &n)
	ss := make([]int, n)
	for i := 0; i < n; i++ {
		scanf("%d ", &ss[i])
	}
	sort.Ints(ss)

	var s, q int
	for _, p := range ss {
		q += p
		s += q
	}
	printf("%d\n", s)
}