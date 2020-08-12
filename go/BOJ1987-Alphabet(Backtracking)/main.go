package main

import (
	"fmt"
	//"os"
)

var R,C int
var slice []string
var isVisit [26]bool
var result int = 0
var dx = [4]int{0, 0, 1, -1}
var dy = [4]int{1, -1, 0, 0}

func max(l int, r int) int {
	if l > r {
		return l
	}
	return r
}

func findSolve(nowY int, nowX int,  depth int) {
	
	if isVisit[slice[nowY][nowX] - 'A'] {
		return
	}
	
	isVisit[slice[nowY][nowX] - 'A'] = true
	result = max(result, depth)
	
	for i := 0; i < 4; i ++ {
		nextY := dy[i] + nowY
		nextX := dx[i] + nowX
		if nextY >= 0 && nextX >= 0 && nextX < C && nextY < R {
			findSolve(nextY, nextX, depth + 1)
		}
	}
	isVisit[slice[nowY][nowX] - 'A'] = false
}

func main() {
	//f, _ := os.Open("input.txt")
	//os.Stdin = f
	
	fmt.Scanf("%d %d", &R, &C)
	
	for i := 0; i < R; i ++ {
		var str string
		fmt.Scanf("%s", &str)
		slice = append(slice, str)
	}
	
	findSolve(0, 0, 1)
	fmt.Printf("%d", result)
}
