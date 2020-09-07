package main

import (
	"container/list"
	"fmt"
)

func solution(progresses []int, speeds []int) []int {

	answer := []int{}
	queue := list.New()
	deployDay := (99-progresses[0])/speeds[0] + 1
	queue.PushBack(0)

	for i := 1; i < len(progresses); i++ {
		day := (99-progresses[i])/speeds[i] + 1
		if day > deployDay {
			deployDay = day
			answer = append(answer, queue.Len())
			for queue.Len() > 0 {
				queue.Remove(queue.Front())
			}
		}
		queue.PushBack(i)
	}
	if queue.Len() > 0 {
		answer = append(answer, queue.Len())
	}

	return answer
}

func main() {
	progresses := [][]int{
		{93, 30, 55},
		{95, 90, 99, 99, 80, 99},
	}
	speeds := [][]int{
		{1, 30, 5},
		{1, 1, 1, 1, 1, 1},
	}

	for i := 0; i < 2; i++ {
		fmt.Println(solution(progresses[i], speeds[i]))
	}

}
