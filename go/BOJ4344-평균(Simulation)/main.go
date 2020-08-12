package main

import (
	"bufio"
	"strconv"
	"os"
	"fmt"
)

func main() {
	r := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	
	r.Scan()
	c, _:= strconv.Atoi(r.Text())
	for i:=0; i < c; i++ {
		r.Scan()
		n, _:= strconv.Atoi(r.Text())
		var arr []int = make([]int, n)
		var c, s int
		for j:=0; j < n; j++ {
			r.Scan()
			item, _:= strconv.Atoi(r.Text())
			arr[j] = item
			s = s + item
		}
		a := float32(s) / float32(n)
		for j:=0; j < n; j++ {
			if float32(arr[j]) > a {
				c++
			}
		}
		fmt.Fprintf(w, "%.3f%%\n", float32(c) / float32(n) * 100)
	}
	w.Flush()
}