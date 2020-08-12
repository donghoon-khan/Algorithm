package main

import(
	"fmt"
	"os"
	"bufio"
)

func main() {
	w := bufio.NewWriter(os.Stdout)

	N := 10000
	self := make([]bool, N + 1)	

	for i := 1; i <= N; i ++ {
		if self[i] {
			continue
		}
		
		next := i
		for {
			tmp := next
			for tmp > 0 {
				next += tmp % 10
				tmp /= 10
			}
			if next > N {
				break
			} else {
				self[next] = true
			}
		}
	}
	
	for i := 1; i <= N; i ++ {
		if !self[i] {
			fmt.Fprintf(w, "%d\n", i)
		}
	}

	w.Flush()
}