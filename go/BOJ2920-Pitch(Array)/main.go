package main

import (
//	"os"
	"fmt"
)

func main(){
//	f, _ := os.Open("input.txt")
//	os.Stdin = f
	
	asc := true
	des := true
	for i := 1; i <= 8; i ++  {
		var num int
		fmt.Scanf("%d", &num)
		
		if (num - i) != 0 {
			asc = false
		}
		if (8 - i - num) != -1 {
			des = false
		}
	}
	if asc {
		fmt.Print("ascending")
	} else if des {
		fmt.Print("descending")
	} else {
		fmt.Print("mixed")
	}
}