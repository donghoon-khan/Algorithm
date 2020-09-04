package main

import (
	"fmt"
	"sort"
	"strings"
)

func solve(phone_book []string) bool {
	sort.Slice(phone_book, func(i, j int) bool {
		return len(phone_book[i]) < len(phone_book[j])
	})
	for i := 0; i < len(phone_book)-1; i++ {
		for j := i + 1; j < len(phone_book); j++ {
			if strings.HasPrefix(phone_book[j], phone_book[i]) {
				return false
			}
		}
	}
	return true
}

func main() {
	phone_book := [][]string{
		{"1195524421", "119", "97674223"},
		{"123", "456", "789"},
		{"12", "123", "1235", "567", "88"},
	}

	for i := 0; i < 3; i++ {
		fmt.Println(solve(phone_book[i]))
	}
}
