package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
	"sort"
)

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	if len(s[i]) < len(s[j]) {
		return true
	} else if len(s[j]) < len(s[i]) {
		return false
	} else if s[i] < s[j] {
		return true
	}
	return false
}

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

	N := NextInt(r)
	str := make([]string, N)
	
	for i := 0; i < N; i ++ {
		str[i] = NextStr(r)
	}
	
	sort.Sort(ByLength(str))
	
	for i := 0; i < N; i ++ {
		if i > 0 {
			if str[i] == str[i - 1] {
				continue
			}
		}
		fmt.Fprintln(w, str[i])
	}
	
	w.Flush()
}