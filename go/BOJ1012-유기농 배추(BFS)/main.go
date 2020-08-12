package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
	"container/list"
)

var (
	r *bufio.Scanner
	w *bufio.Writer
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

func Init() {
	
	w = bufio.NewWriter(os.Stdout)
	r = bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords) 
}

func NextInt() int {
	r.Scan()
	ret, _ := strconv.Atoi(r.Text())

	return ret
}

func NextStr() string {
	r.Scan()
	return r.Text()
}

type Position struct {
	y, x int
}

func main() {
	
	Init()
	
	TC := NextInt()
	
	for t := 0; t < TC; t ++ {
		M := NextInt()
		N := NextInt()
		K := NextInt()
	
		farm := make([][]bool, N)
		for i := 0; i < N; i ++ {
			farm[i] = make([]bool, M)
		}
	
		for i := 0; i < K; i ++ {
			x := NextInt()
			y := NextInt()
			farm[y][x] = true
		}
	
		var result int
		for i := 0; i < N; i ++ {
			for j := 0; j < M; j ++ {
				if farm[i][j] {
					result ++
					q := list.New()
					q.PushBack(Position{i,j})
					farm[i][j] = false
				
					dx := [4]int{1, -1, 0, 0}
					dy := [4]int{0, 0, 1, -1}
					for q.Len() > 0 {
						now := q.Front().Value.(Position)
						q.Remove(q.Front())
					
						for k := 0; k < 4; k ++ {
							var next Position
							next.x = now.x + dx[k]
							next.y = now.y + dy[k]
						
							if next.x >= 0 && next.x < M && next.y >= 0 && next.y < N {
								if farm[next.y][next.x] {
									q.PushBack(next)
									farm[next.y][next.x] = false
								}
							} 
						}
					}
				}
			}
		}
		fmt.Fprintln(w, result)
	}
	w.Flush()
}