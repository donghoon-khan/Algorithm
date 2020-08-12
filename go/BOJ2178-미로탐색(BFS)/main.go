package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
	"container/list"
)

type Position struct {
	x int
	y int
	cnt int
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
	M := NextInt(r)
	
	m := make([]string, N)
	
	for i := 0; i < N; i ++ {
		m[i] = NextStr(r)
	}

	dx := [4]int{1, -1, 0, 0}
	dy := [4]int{0, 0, 1, -1}

	q := list.New()
	visit := make([][]bool, N)
	for i := 0; i < N; i ++ {
		visit[i] = make([]bool, M)
	}
	
	visit[0][0] = true
	q.PushBack(Position{0, 0, 1})
	for q.Len() > 0 {
		now := q.Front().Value.(Position)
		q.Remove(q.Front())
		
		if now.y == N - 1 && now.x == M - 1 {
			fmt.Fprintf(w, "%d\n", now.cnt)
		}
		
		for i := 0; i < 4; i ++ {
			var next Position
			next.x = now.x + dx[i]
			next.y = now.y + dy[i]
			next.cnt = now.cnt + 1
			
			if next.x >= 0 && next.x < M && next.y >= 0 && next.y < N  {
				if m[next.y][next.x] == '1' && !visit[next.y][next.x] {
					visit[next.y][next.x] = true
					q.PushBack(next)
				}
			}
		}
	}
	w.Flush()
}