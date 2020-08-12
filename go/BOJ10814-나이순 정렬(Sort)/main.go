package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
	"sort"
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

type Info struct {
	age int
	name string
	regIdx int
}

func main() {
	
	Init()
	
	N := NextInt()
	m := make([]Info, N)
	
	for i:= 0; i < N; i ++ {
		m[i].age = NextInt()
		m[i].name = NextStr()
		m[i].regIdx = i
	} 
	
	sort.Slice(m, func(i, j int) bool {
			if m[i].age != m[j].age {
				return m[i].age < m[j].age
			} else {
				return m[i].regIdx < m[j].regIdx
			}
	})
	
	for i := 0; i < N; i ++ {
		fmt.Fprintf(w, "%d %s\n", m[i].age, m[i].name)
	}
	
	w.Flush()
}