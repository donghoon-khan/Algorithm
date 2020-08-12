package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var numList []int
var operatorList [4]int
var n, resultMin, resultMax int

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func findSolve(depth, sum int) {
	if depth == n - 1 {
		resultMin = Min(resultMin, sum)
		resultMax = Max(resultMax, sum)
		return
	}
	for i := 0; i < 4; i ++ {
		if operatorList[i] > 0 {
			operatorList[i] --
			if i == 0 {
				findSolve(depth + 1, sum + numList[depth + 1])
			} else if i == 1 {
				findSolve(depth + 1, sum - numList[depth + 1])
			} else if i == 2 {
				findSolve(depth + 1, sum * numList[depth + 1])
			} else if i == 3 {
				findSolve(depth + 1, sum / numList[depth + 1])
			}
			operatorList[i] ++
		}
	}
}

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	defer writer.Flush()

	scanf("%d\n", &n)
	
	numList = make([]int, n)
	for i := 0; i < n; i++ {
		scanf("%d ", &numList[i])
	}
	
	for i := 0; i < 4; i ++ {
		scanf("%d ", &operatorList[i])
	}	
	
	resultMin = math.MaxInt32
	resultMax = math.MinInt32
	
	findSolve(0, numList[0])
	printf("%d\n%d", resultMax, resultMin)
}