package main

import (
  "fmt"
  "math" 
)

func main() {

	var M, N int
	fmt.Scanf("%d", &M)
	fmt.Scanf("%d", &N)
	
	prime := make([]bool, N + 1)
	
	/*--------------------에라스토테네스의 체--------------------*/
	for i := 0; i <= N; i ++ {
		prime[i] = true
	}
	
	prime[0] = false
	prime[1] = false
	for i := 2; float64(i) <= math.Sqrt(float64(N)); i++ {
		if !prime[i] {
			continue
		}
		for j := i + i; j <= N; j += i {
			prime[j] = false
		}
	}
	/*--------------------에라스토테네스의 체--------------------*/
	
	result := 0
	min := math.MaxInt64
	for i := M; i <= N; i ++ {
		if !prime[i] {
			result += i
			if min > i {
				min = i
			}
		}
	}
	
	if result == 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(result)
		fmt.Println(min)
	}
}