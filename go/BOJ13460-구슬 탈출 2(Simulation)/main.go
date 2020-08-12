package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
)

var (
	r *bufio.Scanner
	w *bufio.Writer
	
	board []string
	N, M int
	result int = 11
)

type Position struct {
	x, y int
}

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

func FindNext(red Position, blue Position, x int, y int) (Position, Position, bool) {

	rCnt := 0	
	for board[red.y][red.x] != '#' && board[red.y][red.x] != 'O' {
		red.y += y
		red.x += x
		rCnt ++
	}
	
	bCnt := 0
	for board[blue.y][blue.x] != '#' && board[blue.y][blue.x] != 'O' {
		blue.y += y
		blue.x += x
		bCnt ++
	}
	
	if bCnt > rCnt {
		return red, blue, true
	}
	return red, blue, false
}

func Solve(depth int, prev int, nowR Position, nowB Position) {
	if depth == 10 {
		return
	}
	
	// 0 : R
	// 1 : L
	// 2 : U
	// 3 : D
	dx := [4]int{1, -1, 0, 0}
	dy := [4]int{0, 0, 1, -1}
	for i := 0; i < 4; i ++ {
		if prev == 0 && i == 1 {
			continue
		} else if prev == 1 && i == 0 {
			continue
		} else if prev == 2 && i == 3 {
			continue
		} else if prev == 3 && i == 2 {
			continue
		}
		nextR, nextB, cmp := FindNext(nowR, nowB, dx[i], dy[i])
		
		if board[nextB.y][nextB.x] == 'O' {
			continue
		} else if board[nextR.y][nextR.x] == 'O' {
			result = Min(result, depth + 1)
			return
		} else if nextR.y == nextB.y && nextR.x == nextB.x {
			nextR.y -= dy[i]
			nextB.y -= dy[i]
			nextR.x -= dx[i]
			nextB.x -= dx[i]
			
			if cmp {
				nextB.y -= dy[i]
				nextB.x -= dx[i]
			} else {
				nextR.y -= dy[i]
				nextR.x -= dx[i]
			}
		} else {
			nextR.y -= dy[i]
			nextB.y -= dy[i]
			nextR.x -= dx[i]
			nextB.x -= dx[i]
		}
		
		Solve(depth + 1, i, nextR, nextB)
	}
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

func main() {
	
	Init()
	
	
	N = NextInt()
	M = NextInt()
	
	var red, blue Position
	board = make([]string, N)
	for i := 0; i < N; i ++ {
		
		board[i] = NextStr()
		for j := 0; j < len(board[i]); j ++ {
			if board[i][j] == 'R' {
				red.x = j
				red.y = i
			} else if board[i][j] == 'B' {
				blue.x = j
				blue.y = i
			}
		}	
	}
	
	Solve(0, -1, red, blue)
	
	if result == 11 {
		result = -1
	}
	
	fmt.Fprintln(w, result)
	w.Flush()
}