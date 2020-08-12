package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
	"container/list"
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
	visit := make([]bool, N)
	m := make([][]bool, N)
	for i := 0; i < N; i ++ {
		m[i] = make([]bool, N)
		
	}
	
	M := NextInt(r)
	for i := 0; i < M; i ++ {
		s := NextInt(r)
		e := NextInt(r)
		m[s - 1][e - 1] = true
		m[e - 1][s - 1] = true
	}
	
	q := list.New()
	q.PushBack(0)
	
	result := -1
	visit[0] = true
	for q.Len() > 0 {
		result ++
		now := q.Front().Value.(int)
		
		q.Remove(q.Front())
		
		for i := 0; i < N; i ++ {
			if m[now][i] && !visit[i] {
				visit[i] = true
				q.PushBack(i)
			}
		}
	}
	fmt.Fprintln(w, result)
	w.Flush()
}