package main

import(
	"fmt"
	"os"
	"bufio"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	
	r.Scan()
	str := r.Text()

	var cnt int
	for i := 0; i < len(str); i ++ {
		if str[i] == '=' {
			if i > 1 && str[i - 1] == 'z' && str[i - 2] == 'd' {
				cnt ++
			}
			else if i > 0 {
				if str[i - 1] == 'c' {
					cnt ++
				} else if str[i - 1] == 's' {
					cnt ++
				} else if str[i - 1] == 'z' {
					cnt ++
				}
			} 
			
		} else if str[i] == '-' {
			if i > 0 && str[i - 1] == 'c' {
				cnt ++
			} else if i > 0 && str[i - 1] == 'd' {
				cnt ++
			} 
		} else if str[i] == 'j' {
			if i > 0 && str[i - 1] == 'l' {
				cnt ++
			} else if i > 0 && str[i - 1] == 'n' {
				cnt ++
			}
		}
	}
	
	fmt.Fprintf(w, "%d\n", len(str) - cnt)
	w.Flush()
}