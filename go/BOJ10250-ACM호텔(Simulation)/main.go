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
	T, _ := strconv.Atoi(r.Text())
	
	for i := 0; i < T; i ++ {
		r.Scan()
		H, _ := strconv.Atoi(r.Text())
		r.Scan()
		W, _ := strconv.Atoi(r.Text())
		r.Scan()
		N, _ := strconv.Atoi(r.Text())
		
		var cnt int
		var x, y int
		for x = 1; x <= W; x ++ {
			for y = 1; y <= H; y ++ {
				cnt ++
				if cnt == N {
					break
				}
			}
			if cnt == N {
				break
			}
		}
		
		fmt.Fprintf(w, "%d%02d\n", y, x)	
	}
	

	w.Flush()
}