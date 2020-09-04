package main

import (
	"fmt"
	"sort"
)

func solve(participant []string, completion []string) string {
	sort.Sort(sort.StringSlice(participant))
	sort.Sort(sort.StringSlice(completion))

	for i := 0; i < len(completion); i++ {
		if participant[i] != completion[i] {
			return participant[i]
		}
	}
	return participant[len(participant)-1]
}

func main() {

	participant := [][]string{
		{"leo", "kiki", "eden"},
		{"marina", "josipa", "nikola", "vinko", "filipa"},
		{"mislav", "stanko", "mislav", "ana"},
	}
	completion := [][]string{
		{"eden", "kiki"},
		{"josipa", "filipa", "marina", "nikola"},
		{"stanko", "ana", "mislav"},
	}

	for i := 0; i < 3; i++ {
		fmt.Println(solve(participant[i], completion[i]))
	}
}
