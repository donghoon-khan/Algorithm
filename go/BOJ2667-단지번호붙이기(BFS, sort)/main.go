package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
	"container/list"
	"sort"
)

type Position struct {
	x int
	y int
}

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)
	
	r.Scan()
	N, _ := strconv.Atoi(r.Text())
	
	m := make([]string, N)
	for i := 0; i < N; i ++ {
		r.Scan()
		m[i] = r.Text()
	}
	
	var result []int
	cnt := 0
	dx := [4]int{1, -1, 0, 0}
	dy := [4]int{0, 0, 1, -1}
	visit := make([][]bool, N)
	for i := 0; i < N; i ++ {
		visit[i] = make([]bool, N)
	}
	
	for i := 0; i < N; i ++ {
		for j := 0; j < N; j ++ {
			if m[i][j] == '1' && !visit[i][j] {
				q := list.New()
				cnt ++
				q.PushBack(Position{j, i})
				visit[i][j] = true
				num := 0
				for q.Len() > 0 {
					num ++
					now := q.Front().Value.(Position)
					q.Remove(q.Front())
					
					for k := 0; k < 4; k ++ {
						var next Position
						next.x = now.x + dx[k]
						next.y = now.y + dy[k]
						
						if next.x >= 0 && next.x < N && next.y >= 0 && next.y < N {
							if m[next.y][next.x] == '1' && !visit[next.y][next.x] {
								visit[next.y][next.x] = true
								q.PushBack(next)
							}
						} 
					}
				}
				result = append(result, num)
			}
		}
	}
	sort.Ints(result)
	fmt.Fprintf(w, "%d\n", cnt)
	for i := 0; i < len(result); i ++ {
		fmt.Fprintf(w, "%d\n", result[i])
	}
	w.Flush()
}