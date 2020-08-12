package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

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

	TC := NextInt(r)
	
	for t := 0; t < TC; t ++ {
		str := NextStr(r)
		
		var openCnt, closeCnt int
		result := true
		for i := 0; i < len(str); i ++ {
			if str[i] == '(' {
				openCnt ++
			} else {
				closeCnt ++
			}
			if closeCnt > openCnt {
				result = false
				break
			}
		}
		if openCnt != closeCnt {
			result = false
		}
		if result {
			fmt.Fprintf(w, "%s\n", "YES")
		} else {
			fmt.Fprintf(w, "%s\n", "NO")
		}
	}
	
	w.Flush()
}