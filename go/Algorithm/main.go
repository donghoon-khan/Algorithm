package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"container/list"
)

var reader *bufio.Scanner = bufio.NewScanner(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func NextInt() int {
	reader.Scan()
	ret, _ := strconv.Atoi(reader.Text())
	return ret
}

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }

func main() {
	defer writer.Flush()

	reader.Split(bufio.ScanWords)

	n := NextInt()
	arr := make([]int64, n)
	dp := make([]int64, n)

}