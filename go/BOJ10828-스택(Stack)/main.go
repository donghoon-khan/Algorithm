package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
	"container/list"
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
	stack := list.New()
	for i := 0; i < TC; i ++ {
		cmd := NextStr(r)
		
		if cmd == "push" {
			num := NextInt(r)
			stack.PushBack(num)
		} else if cmd == "pop" {
			ret := -1
			if stack.Len() > 0 {
				ret = stack.Back().Value.(int)
				stack.Remove(stack.Back())
			}
			fmt.Fprintf(w, "%d\n", ret)
		} else if cmd == "size" {
			fmt.Fprintf(w, "%d\n", stack.Len())
		} else if cmd == "empty" {
			ret := 1
			if stack.Len() > 0 {
				ret = 0
			}
			fmt.Fprintf(w, "%d\n", ret)
		} else if cmd == "top" {
			ret := -1
			if stack.Len() > 0 {
				ret = stack.Back().Value.(int)
			}
			fmt.Fprintf(w, "%d\n", ret)
		}
	}
	w.Flush()
}