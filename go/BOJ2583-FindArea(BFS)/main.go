package main

import (
	//"os"
	"fmt"
	"sort"
	"container/list"
)

type Position struct {
	x int
	y int
}

var M, N, K int
var area [100][100]bool
var queue *list.List

func BFS(pos Position) int {
	
	dx := [4]int{0, 0, 1, -1}
	dy := [4]int{1, -1, 0, 0}
	
	ret := 0
	
	area[pos.y][pos.x] = true
	queue.PushBack(pos)
	for queue.Len() > 0 {
		ret ++
		now := queue.Front().Value.(Position)		
		queue.Remove(queue.Front())
		
		for i := 0; i < 4; i ++ {
			var next Position
			next.x = dx[i] + now.x
			next.y = dy[i] + now.y
			
			if next.y >= 0 && next.x >= 0 && next.x < N && next.y < M {
				if !area[next.y][next.x] {
					queue.PushBack(next)
					area[next.y][next.x] = true
				}
			}
		}
	}
	
	return ret
}

func main() {
	//f, _ := os.Open("input.txt")
	//os.Stdin = f
	
	fmt.Scanf("%d %d %d", &M, &N, &K)
	queue = list.New()
	
	for i := 0; i < K; i ++ {
		
		var x1, x2, y1, y2 int
		fmt.Scanf("%d %d %d %d", &x1, &y1, &x2, &y2)
		
		for j := 0; j < y2 - y1; j ++ {
			for k := 0; k < x2 - x1; k ++ {
				area[y1 + j][x1 + k] = true
			}
		}
	}
	
	var result []int
	for i := 0; i < M; i ++ {
		for j := 0; j < N; j ++ {
			if (!area[i][j]) {
				result = append(result, BFS(Position{j,i}))
			}
		}	
	}
	
	sort.Slice(result, func(i, j int) bool {
			return result[i] < result[j]
		})
	
	fmt.Printf("%d\n", len(result))
	for i := 0; i < len(result); i ++ {
		fmt.Printf("%d ", result[i])
	}
}