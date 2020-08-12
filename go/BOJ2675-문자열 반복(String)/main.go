package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)

	r.Scan()
	tc, _ := strconv.Atoi(r.Text())
	
	for i := 0; i < tc; i ++ {
		r.Scan()
		cnt, _ := strconv.Atoi(r.Text())
		
		r.Scan()
		str := r.Text()
		for i := 0; i < len(str); i ++ {
			for j := 0; j < cnt; j ++ {
				fmt.Fprintf(w, "%c", str[i])
			}
		}
		fmt.Fprintf(w, "\n")
	}
	w.Flush()
}