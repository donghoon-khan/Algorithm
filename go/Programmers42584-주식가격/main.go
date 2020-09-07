package main

import "fmt"

func solution(prices []int) []int {

	answer := make([]int, len(prices))

	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[i] > prices[j] {
				answer[i] = j - i
				break
			}
		}
		if answer[i] == 0 {
			answer[i] = len(prices) - 1 - i
		}
	}

	return answer
}

func main() {
	prices := []int{1, 2, 3, 2, 3}

	fmt.Println(solution(prices))
}
