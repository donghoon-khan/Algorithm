package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
)

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

	tc := NextInt(r)
	
	result := 0
	for i := 0; i < tc; i ++ {
		var checker [26]bool
		str := NextStr(r)
		
		checker[str[0] - 'a'] = true
		check := true
		for j := 1; j < len(str); j ++ {
			if !checker[str[j] - 'a'] {
				checker[str[j] - 'a'] = true
			} else if str[j] != str[j - 1] {
				check = false
				break
			}
		}
		if check {
			result ++
		}
	}	
	
	fmt.Fprintf(w, "%d\n", result)
	w.Flush()
}