package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"container/list"
)

var v [][]bool
var w *bufio.Writer

func bfs(N int, start int) {
	q := list.New()
	visit := make([]bool, N + 1)
	
	q.PushBack(start)
	visit[start] = true
	
	for q.Len() > 0 {
		now := q.Front().Value
		q.Remove(q.Front())
		fmt.Fprintf(w, "%d ", now)
		
		for i := 1; i <= N; i ++ {
			if v[now.(int)][i] && !visit[i] {
				visit[i] = true
				q.PushBack(i)
			}
		}
	}
	
}

var visit []bool
func dfs(N int, idx int) {
	visit[idx] = true
	
	fmt.Fprintf(w, "%d ", idx)
	
	for i := 0; i <= N; i ++ {
		if v[idx][i] && !visit[i] {
			dfs(N, i)
		}
	}
}

func main() {
	w = bufio.NewWriter(os.Stdout)
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	
	r.Scan()
	N, _ := strconv.Atoi(r.Text())
	
	r.Scan()
	M, _ := strconv.Atoi(r.Text())
	
	r.Scan()
	start, _ := strconv.Atoi(r.Text())
	
	v = make([][]bool, N + 1)
	visit = make([]bool, N + 1)
	for i := 0; i <= N; i ++ {
		v[i] = make([]bool, N + 1)
	}
	
	for i := 0; i < M; i ++ {
		r.Scan()
		s, _ := strconv.Atoi(r.Text())
		
		r.Scan()
		e, _ := strconv.Atoi(r.Text())
		v[s][e] = true
		v[e][s] = true
	}
	
	dfs(N, start)
	fmt.Fprint(w, "\n")
	bfs(N, start)
	
	w.Flush()
}