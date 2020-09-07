package main

import (
	"container/heap"
	"container/list"
	"fmt"
)

type Element struct {
	priority int
	index    int
}

type PriorityQueue []*Element

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	element := old[n-1]
	element.index = -1
	*pq = old[0 : n-1]
	return element
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	element := x.(*Element)
	element.index = n
	*pq = append(*pq, element)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) update(element *Element, priority int) {
	element.priority = priority
	heap.Fix(pq, element.index)
}

func solution(priorities []int, location int) int {
	answer := 0

	pq := make(PriorityQueue, len(priorities))
	q := list.New()

	for i := range priorities {
		var element Element
		element.index = i
		element.priority = priorities[i]
		pq[i] = &element
		q.PushBack(element)
	}
	heap.Init(&pq)

	for pq.Len() > 0 {

		element := q.Front().Value.(Element)

		if element.priority == pq[0].priority {
			if element.index == location {
				return answer + 1
			} else {
				answer++
				q.Remove(q.Front())
				heap.Pop(&pq)
				//fmt.Println(element.index, element.priority)
			}
		} else {
			q.PushBack(q.Front().Value.(Element))
			q.Remove(q.Front())
		}
	}

	for q.Len() > 0 {
		element := q.Front().Value.(Element)
		fmt.Println(element.priority)
		q.Remove(q.Front())
	}

	return answer
}

func main() {
	priorities := [][]int{
		{2, 1, 3, 2},
		{1, 1, 9, 1, 1, 1},
	}
	location := []int{2, 0}
	for i := 0; i < 2; i++ {
		fmt.Println(solution(priorities[i], location[i]))
	}
}
